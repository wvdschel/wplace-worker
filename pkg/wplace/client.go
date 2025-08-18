// Package wplace provides a client for interacting with the wplace.live API
package wplace

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"io"
	"net/http"
)

const DefaultUserAgent string = "Mozilla/5.0 (X11; Linux x86_64; rv:141.0) Gecko/20100101 Firefox/141.0"
const TileHeight = 1000
const TileWidth = 1000

type Point image.Point

func P(x, y int) Point {
	return Point{
		X: x,
		Y: y,
	}
}

// Client represents a wplace.live API client
type Client struct {
	httpClient *http.Client
	baseURL    string
	userAgent  string
	cookie     string
}

// PixelRequest represents the data needed to paint pixels
type PixelRequest struct {
	Colors []int `json:"colors"`
	Coords []int `json:"coords"`
}

// PixelResponse represents the response from painting pixels
type PixelResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

// UserInfo represents the user information from the /me endpoint
type UserInfo struct {
	AllianceID             int     `json:"allianceId"`
	AllianceRole           string  `json:"allianceRole"`
	Charges                Charges `json:"charges"`
	Country                string  `json:"country"`
	Discord                string  `json:"discord"`
	Droplets               int     `json:"droplets"`
	Email                  string  `json:"email"`
	EquippedFlag           int     `json:"equippedFlag"`
	ExtraColorsBitmap      int     `json:"extraColorsBitmap"`
	FavoriteLocations      []any   `json:"favoriteLocations"`
	FlagsBitmap            string  `json:"flagsBitmap"`
	ID                     int     `json:"id"`
	IsCustomer             bool    `json:"isCustomer"`
	Level                  float64 `json:"level"`
	MaxFavoriteLocations   int     `json:"maxFavoriteLocations"`
	Name                   string  `json:"name"`
	NeedsPhoneVerification bool    `json:"needsPhoneVerification"`
	Picture                string  `json:"picture"`
	PixelsPainted          int     `json:"pixelsPainted"`
	ShowLastPixel          bool    `json:"showLastPixel"`
}
type Charges struct {
	CooldownMs int     `json:"cooldownMs"`
	Count      float64 `json:"count"`
	Max        int     `json:"max"`
}

// NewClient creates a new wplace.live client
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
		baseURL:    "https://backend.wplace.live",
		userAgent:  DefaultUserAgent,
	}
}

func (c *Client) SetCookie(cookie string) {
	c.cookie = cookie
}

func (c *Client) WithCookie(cookie string) *Client {
	c.SetCookie(cookie)
	return c
}

// WithHTTPClient allows setting a custom HTTP client
func (c *Client) WithHTTPClient(client *http.Client) *Client {
	c.httpClient = client
	return c
}

// WithBaseURL allows setting a custom base URL
func (c *Client) WithBaseURL(url string) *Client {
	c.baseURL = url
	return c
}

// PaintPixels paints pixels at the specified coordinates with the specified colors
func (c *Client) PaintPixels(ctx context.Context, tile Point, points []Point, colors []int) (*PixelResponse, error) {
	// Convert points to flat coordinate array
	coords := make([]int, 0, len(points)*2)
	for _, p := range points {
		coords = append(coords, p.X, p.Y)
	}

	// Create the request payload
	payload := PixelRequest{
		Colors: colors,
		Coords: coords,
	}

	// Convert payload to JSON
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	// Create the HTTP request
	url := fmt.Sprintf("%s/s0/pixel/%d/%d", c.baseURL, tile.X, tile.Y)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.7,nl;q=0.3")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("Referer", "https://wplace.live/")
	req.Header.Set("Content-Type", "text/plain;charset=UTF-8")
	req.Header.Set("Origin", "https://wplace.live")
	req.Header.Set("DNT", "1")
	req.Header.Set("Alt-Used", "backend.wplace.live")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", c.cookie)
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("Priority", "u=0")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Cache-Control", "no-cache")

	// Execute the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// Parse the response
	var pixelResp PixelResponse
	if err := json.NewDecoder(resp.Body).Decode(&pixelResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &pixelResp, nil
}

// FetchUserInfo fetches the current user's information from the wplace.live API
func (c *Client) FetchUserInfo(ctx context.Context) (*UserInfo, error) {
	// Create the HTTP request
	url := fmt.Sprintf("%s/me", c.baseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers based on the curl request
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.7,nl;q=0.3")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("Referer", "https://wplace.live/")
	req.Header.Set("Origin", "https://wplace.live")
	req.Header.Set("DNT", "1")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", c.cookie)

	// Execute the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	body := respBody(resp)
	defer body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		data, _ := io.ReadAll(body)
		return nil, fmt.Errorf("unexpected status code: %d: %s", resp.StatusCode, string(data))
	}

	// Parse the response
	var userInfo UserInfo
	if err := json.NewDecoder(body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &userInfo, nil
}

// FetchTile fetches an image tile from the wplace.live API
func (c *Client) FetchTile(ctx context.Context, server int, tile Point) (image.Image, error) {
	// Construct the URL based on the pattern from the cURL request
	url := fmt.Sprintf("%s/files/s%d/tiles/%d/%d.png", c.baseURL, server, tile.X, tile.Y)

	// Create the HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers based on the cURL request
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Accept", "image/webp,*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.7,nl;q=0.3")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("Referer", "https://wplace.live/")
	req.Header.Set("Origin", "https://wplace.live")
	req.Header.Set("DNT", "1")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("Connection", "keep-alive")

	// Add cookie if available
	if c.cookie != "" {
		req.Header.Set("Cookie", c.cookie)
	}

	// Execute the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		body := respBody(resp)
		data, err := io.ReadAll(body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %w", err)
		}

		return nil, fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, string(data))
	}

	// Determine content type from response headers
	contentType := resp.Header.Get("Content-Type")
	if contentType != "image/png" {
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %w", err)
		}

		return nil, fmt.Errorf("unexpected content type code: %s, response: %s", contentType, string(data))
	}

	img, err := png.Decode(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return img, nil
}

func (c *Client) FetchImage(ctx context.Context, tile Point, pixel Point, dimensions Point) (image.Image, error) {
	res := image.NewRGBA(image.Rectangle{
		Max: image.Point(dimensions),
	})

	currentTile := tile
	currentPixel := pixel
	posInRes := P(0, 0)

	for posInRes.Y < dimensions.Y {
		tileImg, err := c.FetchTile(ctx, 0, currentTile)
		if err != nil {
			return nil, err
		}
		posInTile := P(currentPixel.X, currentPixel.Y)
		widthOfTileSection := min(dimensions.X-posInRes.X, tileImg.Bounds().Max.X)
		heightOfTileSection := min(dimensions.Y-posInRes.Y, tileImg.Bounds().Max.Y)

		blitImage(tileImg, res,
			image.Rect(posInTile.X, posInTile.Y, widthOfTileSection, heightOfTileSection),
			image.Point(posInRes))

		posInRes.X += widthOfTileSection
		currentTile.X += 1
		if posInRes.X >= dimensions.X {
			posInRes.X = 0
			posInRes.Y += heightOfTileSection
			currentTile.Y += 1
			currentTile.X = tile.X
		}
	}

	return res, nil
}

func min(x ...int) int {
	minX := x[0]
	for _, i := range x {
		if i < minX {
			minX = i
		}
	}
	return minX
}
