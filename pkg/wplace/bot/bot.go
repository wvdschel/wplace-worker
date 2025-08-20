package bot

import (
	"context"
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
		cookies:      cookies,
		lock:         &sync.Mutex{},
		accounts:     make([]wplace.UserInfo, len(config.Cookies)),
		images:       imgStatus,
	}, nil
}

func (b *Bot) update(ctx context.Context) error {
	for i, img := range b.images {
		var current image.Image
		img.updated = false

		err := retry.Do(func() error {
			var err error
			current, err = b.wplaceClient.FetchImage(ctx, img.tile, img.pixel, wplace.P(img.target.Bounds().Dx(), img.target.Bounds().Dy()))
			return err
		}, defaultRetryOpts...)

		if err != nil {
			log.Printf("Error fetching image #%d: %v\n", i+1, err)
			continue
		}
		img.update(current)
	}

	for i, img := range b.images {
		log.Printf("Image #%d: %d of %d pixels done, %d to go\n", i+1, img.correctPixelCount, img.totalPixelCount, img.totalPixelCount-img.correctPixelCount)
		if img.current != nil {
			kittyimg.Fprintln(os.Stdout, img.getImage())
		}
	}

	for i := range b.cookies {
		if err := b.refreshCloudFlareToken(ctx, i, false); err != nil {
			log.Printf("Error refreshing CloudFlare token for account %d: %v\n", i+1, err)
		}
	}

	for i, c := range b.cookies {
		b.wplaceClient.SetCookies(c)

		var accInfo *wplace.UserInfo
		err := retry.Do(func() error {
			var err error
			accInfo, err = b.wplaceClient.FetchUserInfo(ctx)
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

func (b *Bot) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Run one update before running the painters
	if err := b.update(ctx); err != nil {
		return err
	}

	for i := range b.cookies {
		go b.painter(ctx, i)
	}

	for {
		if err := b.update(ctx); err != nil {
			return err
		}

		time.Sleep(time.Second * 3)
	}
}
