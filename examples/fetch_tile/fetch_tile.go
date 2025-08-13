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
	client := wplace.NewClient("")

	resp, err := client.FetchTile(context.Background(), 0, wplace.P(1222, 832))
	if err != nil {
		log.Fatal(err)
	}

	kittyimg.Fprintln(os.Stdout, resp)
}
