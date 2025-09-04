package bot

import (
	"bytes"
	"embed"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"log"
	"net/http"
	"time"
)

//go:embed dashboard.html
var dashboardHTML embed.FS

// AccountStatus represents the status information for an account
type AccountStatus struct {
	Name          string  `json:"name"`
	ChargeCount   float64 `json:"chargeCount"`
	ChargeMax     int     `json:"chargeMax"`
	ChargePercent float64 `json:"chargePercent"`
	OverflowTime  string  `json:"overflowTime"`
	Country       string  `json:"country"`
	Level         float64 `json:"level"`
	PixelsPainted int     `json:"pixelsPainted"`
}

// ImageProgress represents the progress information for an image
type ImageProgress struct {
	Index           int    `json:"index"`
	TotalPixels     int    `json:"totalPixels"`
	CorrectPixels   int    `json:"correctPixels"`
	IncorrectPixels int    `json:"incorrectPixels"`
	ProgressPercent int    `json:"progressPercent"`
	ImageData       string `json:"imageData"`
}

// DashboardData represents all the data needed for the dashboard
type DashboardData struct {
	Accounts      []AccountStatus `json:"accounts"`
	Images        []ImageProgress `json:"images"`
	Logs          []LogEntry      `json:"logs"`
	TotalCharges  int             `json:"totalCharges"`
	TotalCapacity int             `json:"totalCapacity"`
}

// encodeImageToPNG encodes an image to PNG format and returns the bytes
func encodeImageToPNG(img image.Image) ([]byte, error) {
	// Create a buffer to hold the PNG data
	buf := new(bytes.Buffer)
	// Encode the image to PNG format
	err := png.Encode(buf, img)
	if err != nil {
		return nil, err
	}
	// Return the bytes
	return buf.Bytes(), nil
}

// WebHandler handles all web server requests
type WebHandler struct {
	bot *Bot
}

func NewWebHandler(bot *Bot) *WebHandler {
	return &WebHandler{bot: bot}
}

// handleDashboard serves the main dashboard page
func (h *WebHandler) handleDashboard(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Read the embedded HTML file
	htmlFile, err := dashboardHTML.ReadFile("dashboard.html")
	if err != nil {
		http.Error(w, "Failed to load dashboard", http.StatusInternalServerError)
		return
	}

	w.Write(htmlFile)
}

// handleAPI serves the API endpoint for dashboard data
func (h *WebHandler) handleAPI(w http.ResponseWriter, r *http.Request) {
	images := make([]image.Image, len(h.bot.images))
	for i, img := range h.bot.images {
		if img.current != nil {
			images[i] = img.getImage()
		}
	}

	h.bot.lock.RLock()
	defer h.bot.lock.RUnlock()

	data := DashboardData{
		Accounts:      make([]AccountStatus, len(h.bot.accounts)),
		Images:        make([]ImageProgress, len(h.bot.images)),
		Logs:          h.bot.logBuffer,
		TotalCharges:  0,
		TotalCapacity: 0,
	}

	// Process account data
	for i, acc := range h.bot.accounts {
		userInfo := acc.userInfo
		chargePercent := 0.0
		if userInfo.Charges.Max > 0 {
			chargePercent = (userInfo.Charges.Count / float64(userInfo.Charges.Max)) * 100
		}

		capacityLeft := float64(userInfo.Charges.Max) - userInfo.Charges.Count
		timeUntilOverflow := time.Second * time.Duration(30*capacityLeft)
		overflowTimestamp := time.Now().Add(timeUntilOverflow)

		data.Accounts[i] = AccountStatus{
			Name:          userInfo.Name,
			ChargeCount:   userInfo.Charges.Count,
			ChargeMax:     userInfo.Charges.Max,
			ChargePercent: chargePercent,
			OverflowTime:  overflowTimestamp.Format("15:04"),
			Country:       userInfo.Country,
			Level:         userInfo.Level,
			PixelsPainted: userInfo.PixelsPainted,
		}

		data.TotalCharges += int(userInfo.Charges.Count)
		data.TotalCapacity += userInfo.Charges.Max
	}

	// Process image data
	for i, img := range h.bot.images {
		incorrectPixels := img.totalPixelCount - img.correctPixelCount
		progressPercent := 0
		if img.totalPixelCount > 0 {
			progressPercent = (img.correctPixelCount * 100) / img.totalPixelCount
		}

		// Get the image data using getImage() and encode as base64 PNG
		var imageData string
		if img.current != nil {
			buf, err := encodeImageToPNG(images[i])
			if err != nil {
				log.Printf("Error encoding image %d: %v", i, err)
			} else {
				imageData = base64.StdEncoding.EncodeToString(buf)
			}
		}

		data.Images[i] = ImageProgress{
			Index:           i,
			TotalPixels:     img.totalPixelCount,
			CorrectPixels:   img.correctPixelCount,
			IncorrectPixels: incorrectPixels,
			ProgressPercent: progressPercent,
			ImageData:       imageData,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// StartWebServer starts the web server
func (b *Bot) StartWebServer(port int) error {
	handler := NewWebHandler(b)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.handleDashboard)
	mux.HandleFunc("/api/dashboard", handler.handleAPI)

	b.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	go func() {
		if err := b.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Web server error: %v", err)
		}
	}()

	return nil
}

// StopWebServer stops the web server
func (b *Bot) StopWebServer() error {
	if b.server != nil {
		return b.server.Close()
	}
	return nil
}
