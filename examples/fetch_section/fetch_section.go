package main

import (
	"context"
	"log"
	"os"

	"github.com/dolmen-go/kittyimg"
	"github.com/wvdschel/wplace-worker/pkg/wplace"
)

func main() {
	// Create a new client with your session cookie
	client := wplace.NewClient()

	resp, err := client.FetchImage(context.Background(), wplace.P(1221, 831), wplace.P(970, 990), wplace.P(512, 512))
	if err != nil {
		log.Fatal(err)
	}

	kittyimg.Fprintln(os.Stdout, resp)

	resp, err = client.FetchImage(context.Background(), wplace.P(1044, 685), wplace.P(134, 434), wplace.P(371, 305))
	if err != nil {
		log.Fatal(err)
	}

	kittyimg.Fprintln(os.Stdout, resp)

	resp, err = client.FetchImage(context.Background(), wplace.P(926, 708), wplace.P(0, 0), wplace.P(32, 1))
	if err != nil {
		log.Fatal(err)
	}

	kittyimg.Fprintln(os.Stdout, resp)
}
