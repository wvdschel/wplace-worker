package bot

import (
	"sync"
	"time"

	"github.com/jrsap/wplace-worker/pkg/wplace"
)

type Config struct {
	Cookies []string `json:"cookies"`

	Templates []Template `json:"templates"`

	Limits Limits `json:"limits"`
}

type Template struct {
	Tile  wplace.Point `json:"tile"`
	Pixel wplace.Point `json:"pixel"`
	Path  string       `json:"path"`
}

type Limits struct {
	MaxPixelsPerRequest   int            `json:"maxPixelsPerRequest,omitempty"`
	MinDelayBetweenPaints *time.Duration `json:"minDelayBetweenPaints,omitempty"`
	MaxDelayBetweenPaints *time.Duration `json:"maxDelayBetweenPaints,omitempty"`
}

type Bot struct {
	config *Config

	client *wplace.Client // client for each cookie

	lock *sync.Mutex

	log      chan string
	userInfo chan wplace.UserInfo
	images   []imageStatus
}

func New(config *Config) (*Bot, error) {
	c := wplace.NewClient()

	imgStatus := make([]imageStatus, len(config.Templates))

	for idx, tmp := range config.Templates {
		i, err := wplace.LoadImage(tmp.Path)
		if err != nil {
			return nil, err
		}
		imgStatus[idx].current = wplace.ConvertToPallette(i)
	}

	return &Bot{
		config:   config,
		client:   c,
		lock:     &sync.Mutex{},
		log:      make(chan string, 4),
		userInfo: make(chan wplace.UserInfo, len(config.Cookies)),
		images:   imgStatus,
	}, nil
}

func (b *Bot) getNextPixels() (tile wplace.Point, pixels []wplace.Point) {
	b.lock.Lock()
	defer b.lock.Unlock()

	// TODO fetch images and compare to the templates
	return wplace.Point{}, []wplace.Point{}
}

func (b *Bot) cancelPixels(tile wplace.Point, pixels []wplace.Point) {
	b.lock.Lock()
	defer b.lock.Unlock()
}

func (b *Bot) update() error {
	b.lock.Lock()
	defer b.lock.Unlock()

	// TODO for all templates: print target image, current state with highlighted mismatched pixels, and pixels done vs pixels total
	// TODO

	return nil
}

func (b *Bot) Run() error {
	return b.update()
}

func ExampleConfig() *Config {
	min := 30 * time.Second
	max := 5 * time.Minute

	return &Config{
		Cookies: []string{"cookie=1234567890"},
		Templates: []Template{
			{
				Tile:  wplace.P(1, 1),
				Pixel: wplace.P(1, 1),
				Path:  "a.png",
			},
		},
		Limits: Limits{
			MaxPixelsPerRequest:   20,
			MinDelayBetweenPaints: &min,
			MaxDelayBetweenPaints: &max,
		},
	}
}
