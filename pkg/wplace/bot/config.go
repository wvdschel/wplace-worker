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
	MaxPixelsPerRequest     int `json:"maxPixelsPerRequest,omitempty"`
	MinSecondsBetweenPaints int `json:"minSecondsBetweenPaints,omitempty"`
	MaxSecondsBetweenPaints int `json:"maxSecondsBetweenPaints,omitempty"`
}

func ExampleConfig() *Config {
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
			MaxPixelsPerRequest:     20,
			MinSecondsBetweenPaints: 30,
			MaxSecondsBetweenPaints: 300,
		},
		Byparr: &ByparrConfig{
			BaseURL:     "http://localhost:8191/v1",
			MaxRetries:  &retries,
			MaxDuration: &duration,
		},
	}
}
