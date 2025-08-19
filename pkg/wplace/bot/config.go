package bot

import (
	"time"

	"github.com/jrsap/wplace-worker/pkg/wplace"
)

type Config struct {
	Cookies []string `json:"cookies"`

	Templates []Template `json:"templates"`

	Limits Limits `json:"limits"`

	Byparr *ByparrConfig `json:"byparr,omitempty"`
}

type ByparrConfig struct {
	BaseURL     string         `json:"baseURL,omitempty"`
	MaxRetries  *int           `json:"maxRetries,omitempty"`
	MaxDuration *time.Duration `json:"maxDuration,omitempty"`
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

func ExampleConfig() *Config {
	min := 30 * time.Second
	max := 5 * time.Minute

	retries := 4
	duration := 15 * time.Second

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
		Byparr: &ByparrConfig{
			BaseURL:     "http://localhost:8191/v1",
			MaxRetries:  &retries,
			MaxDuration: &duration,
		},
	}
}
