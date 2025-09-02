package bot

import (
	"net/http"

	"github.com/jrsap/wplace-worker/pkg/wplace"
)

type Config struct {
	Accounts    []Account         `json:"accounts"`
	Templates   []Template        `json:"templates"`
	Limits      Limits            `json:"limits"`
	CloudBuster CloudBusterConfig `json:"cloudbuster"`
	WebPort     int               `json:"webPort,omitempty"`
}

type Template struct {
	Tile     wplace.Point `json:"tile"`
	Pixel    wplace.Point `json:"pixel"`
	Path     string       `json:"path"`
	Disabled bool         `json:"disabled,omitempty"`
}

type Limits struct {
	MaxPixelsPerRequest     int `json:"maxPixelsPerRequest,omitempty"`
	MinPixelsPerRequest     int `json:"minPixelsPerRequest,omitempty"`
	MinSecondsBetweenPaints int `json:"minSecondsBetweenPaints,omitempty"`
}

func ExampleConfig() *Config {
	return &Config{
		Accounts: []Account{{
			Cookie: "j=1234567890",
		}},
		Templates: []Template{
			{
				Tile:  wplace.P(1, 1),
				Pixel: wplace.P(1, 1),
				Path:  "a.png",
			},
		},
		Limits: Limits{
			MaxPixelsPerRequest:     100,
			MinPixelsPerRequest:     20,
			MinSecondsBetweenPaints: 30,
		},
		CloudBuster: CloudBusterConfig{
			BaseURL:    "http://localhost:8000",
			MaxRetries: 3,
		},
		WebPort: 8080,
	}
}

type Account struct {
	Cookie   string `json:"cookie"`
	ReadOnly bool   `json:"readonly,omitempty"`
	Jumphost string `json:"jumphost,omitempty"`

	userInfo wplace.UserInfo
	cookies  []*http.Cookie
}

type CloudBusterConfig struct {
	BaseURL    string `json:"baseURL"`
	MaxRetries uint   `json:"maxRetries"`
}
