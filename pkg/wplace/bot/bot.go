package bot

import (
	"context"
	"image"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/avast/retry-go"
	"github.com/wvdschel/wplace-worker/pkg/cloudbuster"
	"github.com/wvdschel/wplace-worker/pkg/wplace"
)

type Bot struct {
	config       *Config
	wplaceClient *wplace.Client
	cloudbuster  *cloudbuster.Client

	images   []*imageStatus
	accounts []*Account

	// Web server fields
	logBuffer []LogEntry
	logLock   *sync.RWMutex
	server    *http.Server
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

	accs := make([]*Account, len(config.Accounts))
	for i, acc := range config.Accounts {
		accs[i] = &acc
		accs[i].lock = &sync.Mutex{}
		accs[i].client = wplace.NewClient()
		var err error
		accs[i].cookies, err = http.ParseCookie(acc.Cookie)
		if err != nil {
			return nil, err
		}
	}

	return &Bot{
		config:       config,
		wplaceClient: c,
		accounts:     accs,
		images:       imgStatus,
		logBuffer:    make([]LogEntry, 0),
		logLock:      &sync.RWMutex{},
		cloudbuster:  cloudbuster.NewClient(config.CloudBuster.BaseURL, http.DefaultClient),
	}, nil
}

func (b *Bot) updateUserInfo(ctx context.Context, i int) error {
	acc := b.accounts[i]
	acc.lock.Lock()
	defer acc.lock.Unlock()

	acc.client.SetCookies(acc.cookies)

	var userInfo *wplace.UserInfo
	err := retry.Do(func() error {
		var err error
		userInfo, err = acc.client.FetchUserInfo(ctx)
		return err
	}, defaultRetryOpts...)

	if err != nil {
		return err
	}
	b.accounts[i].userInfo = *userInfo

	return nil
}

func (b *Bot) refreshImages(ctx context.Context) {
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
}

func (b *Bot) update(ctx context.Context) error {
	b.refreshImages(ctx)

	for i, img := range b.images {
		log.Printf("Image #%d: %d of %d pixels done, %d to go\n", i+1, img.correctPixelCount, img.totalPixelCount, img.totalPixelCount-img.correctPixelCount)
		// if img.current != nil {
		// 	kittyimg.Fprintln(os.Stdout, img.getImage())
		// }
	}

	for i := range b.accounts {
		if err := b.updateUserInfo(ctx, i); err != nil {
			log.Printf("Error getting account info for account %d: %v\n", i+1, err)
		}
	}

	totalCharges, totalCapacity := 0, 0
	for i, acc := range b.accounts {
		userInfo := acc.userInfo
		capacityLeft := float64(userInfo.Charges.Max) - userInfo.Charges.Count
		timeUntilOverflow := time.Second * time.Duration(30*capacityLeft)
		overflowTimestamp := time.Now().Add(timeUntilOverflow)

		log.Printf("#%d: %s has %0.0f/%d pixels - first overflow at %s\n", i+1, userInfo.Name, userInfo.Charges.Count, userInfo.Charges.Max, overflowTimestamp.Format("15:04"))
		totalCapacity += userInfo.Charges.Max
		totalCharges += int(userInfo.Charges.Count)
	}

	log.Printf("Total: %d/%d pixels\n", totalCharges, totalCapacity)

	return nil
}

func (b *Bot) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Start web server if port is configured
	if b.config.WebPort > 0 {
		if err := b.StartWebServer(b.config.WebPort); err != nil {
			log.Printf("Failed to start web server: %v", err)
		} else {
			log.Printf("Web server started on port %d", b.config.WebPort)
		}
	}

	// Run one update before running the painters
	if err := b.update(ctx); err != nil {
		return err
	}

	for i, acc := range b.accounts {
		if !acc.ReadOnly {
			go b.painter(ctx, i)
		}
	}

	// Wait for context cancellation or error
	done := make(chan error, 1)
	go func() {
		for {
			time.Sleep(time.Second * 30)
			if err := b.update(ctx); err != nil {
				done <- err
				return
			}
		}
	}()

	select {
	case err := <-done:
		// Update error occurred
		log.Printf("Bot stopped due to error: %v", err)
		return err
	case <-ctx.Done():
		// Context was cancelled (graceful shutdown)
		log.Println("Bot shutting down gracefully...")
		return b.StopWebServer()
	}
}
