package main

import (
	"context"
	"flag"
	"fmt"
	"image/color"
	"image/png"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dolmen-go/kittyimg"
	"github.com/jrsap/wplace-worker/pkg/cloudbuster"
	"github.com/jrsap/wplace-worker/pkg/wplace"
)

var (
	cloudbusterURL = flag.String("cloudbuster", "http://localhost:8000", "Cloudbuster URL")
	tx             = flag.Int("X", 926, "Tile X coordinate")
	ty             = flag.Int("Y", 708, "Tile Y coordinate")
)

func main() {
	flag.Parse()

	// Get cookie from environment variable
	cookie := os.Getenv("WPLACE_COOKIE")
	if cookie == "" {
		log.Fatal("WPLACE_COOKIE environment variable is required")
	}

	cookies, err := http.ParseCookie(cookie)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new client with your session cookie
	client := wplace.NewClient()
	cb := cloudbuster.NewClient(*cloudbusterURL, http.DefaultClient)

	token, cf_cookies, err := cb.GetToken("https://wplace.live", "")
	if err != nil {
		log.Fatal(err)
	}

	cookies = append(cookies, cf_cookies...)

	tile := wplace.P(*tx, *ty)
	pixels := make([]wplace.Point, 32)
	colors := make([]int, 32)
	for i := range pixels {
		pixels[i].X = i
		colors[i] = i
	}

	client.SetCookies(cookies)
	_, err = client.PaintPixels(context.Background(), token, tile, pixels, colors)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(10 * time.Second)

	img, err := client.FetchImage(context.Background(), tile, wplace.P(0, 0), wplace.P(32, 1))
	if err != nil {
		log.Fatal(err)
	}

	kittyimg.Fprintln(os.Stdout, img)

	f, err := os.OpenFile("palette.png", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 32; i++ {
		fmt.Printf("ConvertColor(\"%s\"),\n", colorToHex(img.At(i, 0)))
	}
}

func colorToHex(c color.Color) string {
	r, g, b, a := c.RGBA()
	return fmt.Sprintf("#%02x%02x%02x%02x", r>>8, g>>8, b>>8, a>>8)
}
