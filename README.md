# wplace-worker

A Go client library for interacting with the wplace.live API.

## Installation

```bash
go get github.com/wim/wplace-worker
```

## Usage

```go
package main

import (
    "context"
    "fmt"
    "log"
    "github.com/wim/wplace-worker/pkg/wplace"
)

func main() {
    // Create a new client with your session cookie
    cookie := "cf_clearance=UmEFo_mCRC1TeC4XYn8Oh_dx52VViqcfO9eIe3theyQ-1754657418-1.2.1.1-mdE6vkXOZ3YgaGQAOjOCmalf4Bc1XoKNDSw67fbnA31od6n9ITmB7zvg24MhZYpn0p1k6EbLKLoJw_dBxF19defVadpQXc4U6girNblyGQ_XfXALDjdZknoC1TLl9tZIvAEKFNAOCa5rJ.5mYC6sTN54BRmmxbGQMeW.bWG7kutaMATaU4bS8UGWcrA3kWIhfhUvV1WLvIe74UcsQ2VhNw8PJ1gvtzUoI7UbYFVvtWo; s=grnRccQ1cqp3oaoDumE1kQ%3D%3D"
    client := wplace.NewClient(cookie)
    
    // Paint pixels at coordinates (459,167) and (460,168) with color 8
    resp, err := client.PaintPixels(context.Background(), 459, 167, 460, 168, 8, 8)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Success: %v, Message: %s\n", resp.Success, resp.Message)
}
```

## API Reference

### NewClient

```go
func NewClient(cookie string) *Client
```

Creates a new wplace.live client with the provided session cookie.

### PaintPixels

```go
func (c *Client) PaintPixels(ctx context.Context, x1, y1, x2, y2 int, color1, color2 int) (*PixelResponse, error)
```

Paints pixels at the specified coordinates with the specified colors.

Parameters:
- `ctx`: Context for the request
- `x1, y1`: First pixel coordinates
- `x2, y2`: Second pixel coordinates
- `color1`: Color for the first pixel
- `color2`: Color for the second pixel

Returns:
- `PixelResponse`: Response from the API
- `error`: Any error that occurred

### WithHTTPClient

```go
func (c *Client) WithHTTPClient(client *http.Client) *Client
```

Allows setting a custom HTTP client.

### WithBaseURL

```go
func (c *Client) WithBaseURL(url string) *Client
```

Allows setting a custom base URL.