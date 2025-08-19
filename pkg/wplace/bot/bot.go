package bot

import (
	"context"
	"fmt"
	"image"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/avast/retry-go"
	"github.com/dolmen-go/kittyimg"
	"github.com/jrsap/wplace-worker/pkg/byparr"
	"github.com/jrsap/wplace-worker/pkg/wplace"
)

type Bot struct {
	config *Config

	wplaceClient *wplace.Client
	byparrClient *byparr.Client

	lock *sync.Mutex

	images   []*imageStatus
	cookies  [][]*http.Cookie
	accounts []wplace.UserInfo
}

func New(config *Config) (*Bot, error) {
	c := wplace.NewClient()

	imgStatus := make([]*imageStatus, len(config.Templates))

	for idx, tmp := range config.Templates {
		i, err := wplace.LoadImage(tmp.Path)
		if err != nil {
			return nil, err
		}

		imgStatus[idx] = newImageStatus(tmp.Tile, tmp.Pixel, wplace.ConvertToPallette(i))
	}

	cookies := make([][]*http.Cookie, len(config.Cookies))
	for idx, cookieString := range config.Cookies {
		var err error
		cookies[idx], err = http.ParseCookie(cookieString)
		if err != nil {
			return nil, err
		}
	}

	var bc *byparr.Client
	if config.Byparr != nil {
		opts := []byparr.Option{}
		if config.Byparr.BaseURL != "" {
			opts = append(opts, byparr.WithBaseURL(config.Byparr.BaseURL))
		}

		c, err := byparr.New(opts...)
		if err != nil {
			return nil, err
		}
		bc = c
	}

	return &Bot{
		config:       config,
		wplaceClient: c,
		byparrClient: bc,
		lock:         &sync.Mutex{},
		accounts:     make([]wplace.UserInfo, len(config.Cookies)),
		images:       imgStatus,
	}, nil
}

func (b *Bot) getNextPixels() (tile wplace.Point, pixels []wplace.Point) {
	b.lock.Lock()
	defer b.lock.Unlock()

	for _, img := range b.images {
		tile, pixels = img.getWork(b.config.Limits.MaxPixelsPerRequest)

		if len(pixels) != 0 {
			return tile, pixels
		}
	}

	return wplace.P(0, 0), nil
}

func (b *Bot) cancelPixels(tile wplace.Point, pixels []wplace.Point) {
	b.lock.Lock()
	defer b.lock.Unlock()

	if len(pixels) == 0 {
		return
	}

	for _, img := range b.images {
		if img.contains(tile, pixels[0]) {
			img.returnWork(tile, pixels)
			return
		}
	}
}

func (b *Bot) update() error {
	b.lock.Lock()
	defer b.lock.Unlock()

	// TODO for all templates: print target image, current state with highlighted mismatched pixels, and pixels done vs pixels total
	for i, img := range b.images {
		var current image.Image

		log.Printf("Fetching image #%d\n", i+1)
		err := retry.Do(func() error {
			var err error
			current, err = b.wplaceClient.FetchImage(context.Background(), img.tile, img.pixel, wplace.P(img.target.Bounds().Dx(), img.target.Bounds().Dy()))
			return err
		}, defaultRetryOpts...)

		if err != nil {
			log.Printf("Error fetching image #%d: %v\n", i+1, err)
			continue
		}
		img.update(current)
	}

	for i, img := range b.images {
		log.Printf("Image #%d: %d of %d pixels done\n", i+1, img.correctPixelCount, img.totalPixelCount)
		if img.current != nil {
			kittyimg.Fprintln(os.Stdout, img.getImage())
		}
		fmt.Println()
	}

	for i, c := range b.cookies {
		cfClearance, err := b.withCFClearance(context.Background(), c)
		if err != nil {
			log.Printf("Error getting CF token for account #%d: %v\n", i+1, err)
			continue
		}
		b.wplaceClient.SetCookies(cfClearance)

		log.Printf("Fetching account info #%d\n", i+1)
		var accInfo *wplace.UserInfo
		err = retry.Do(func() error {
			var err error
			accInfo, err = b.wplaceClient.FetchUserInfo(context.Background())
			return err
		}, defaultRetryOpts...)

		if err != nil {
			log.Printf("Error getting account info for account %d: %v\n", i+1, err)
			continue
		}
		b.accounts[i] = *accInfo
	}

	for i, accInfo := range b.accounts {
		log.Printf("#%d: %s has %0.0f/%d pixels\n", i+1, accInfo.Name, accInfo.Charges.Count, accInfo.Charges.Max)
	}

	return nil
}

func (b *Bot) Run() error {
	for {
		if err := b.update(); err != nil {
			return err
		}

		// TODO draw pixels
		time.Sleep(time.Second * 3)
	}
}
