package main

import (
	"context"
	"log"
	"os"

	"github.com/dolmen-go/kittyimg"
	"github.com/jrsap/wplace-worker/pkg/wplace"
)

func main() {
	// Create a new client with your session cookie
	client := wplace.NewClient()

	resp, err := client.FetchImage(context.Background(), wplace.P(1221, 832), wplace.P(0, 0), wplace.P(1900, 1200))
	if err != nil {
		log.Fatal(err)
	}

	kittyimg.Fprintln(os.Stdout, resp)
}
