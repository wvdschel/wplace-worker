// Package wplace provides a client for interacting with the wplace.live API
package wplace

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const DefaultUserAgent string = "Mozilla/5.0 (X11; Linux x86_64; rv:141.0) Gecko/20100101 Firefox/141.0"

type Point struct {
	X, Y int
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

// ImageResponse represents the response from fetching an image
type ImageResponse struct {
	Data        []byte
	ContentType string
}

// NewClient creates a new wplace.live client
func NewClient(cookie string) *Client {
	return &Client{
		httpClient: &http.Client{},
		baseURL:    "https://backend.wplace.live",
		userAgent:  DefaultUserAgent,
		cookie:     cookie,
	}
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
func (c *Client) PaintPixels(ctx context.Context, points []Point, colors []int) (*PixelResponse, error) {
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
	url := fmt.Sprintf("%s/s0/pixel/%d/%d", c.baseURL, 1222, 832)
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

// FetchTile fetches an image tile from the wplace.live API
func (c *Client) FetchTile(ctx context.Context, server, tileX, tileY int) (*ImageResponse, error) {
	// Construct the URL based on the pattern from the cURL request
	url := fmt.Sprintf("%s/files/s%d/tiles/%d/%d.png", c.baseURL, server, tileX, tileY)

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
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, string(data))
	}

	// Determine content type from response headers
	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		// Default to image/png if not specified
		contentType = "image/png"
	}

	return &ImageResponse{
		Data:        data,
		ContentType: contentType,
	}, nil
}
