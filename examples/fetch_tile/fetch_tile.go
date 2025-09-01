package main

import (
	"context"
	"flag"
	"fmt"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/dolmen-go/kittyimg"
	"github.com/jrsap/wplace-worker/pkg/wplace"
)

var (
	x = flag.Int("x", 1045, "x coordinate of the tile")
	y = flag.Int("y", 685, "y coordinate of the tile")
)

func main() {
	flag.Parse()

	// Create a new client with your session cookie
	client := wplace.NewClient()

	resp, err := client.FetchTile(context.Background(), 0, wplace.P(*x, *y))
	if err != nil {
		log.Fatal(err)
	}

	kittyimg.Fprintln(os.Stdout, resp)

	f, err := os.OpenFile(fmt.Sprintf("%d_%d_%s.png", *x, *y, time.Now().Format(time.RFC3339)), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if err := png.Encode(f, resp); err != nil {
		log.Fatal(err)
	}
}
