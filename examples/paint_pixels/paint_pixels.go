// Package main demonstrates how to use the wplace client library
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/wvdschel/wplace-worker/pkg/wplace"
)

func main() {
	// Get cookie from environment variable
	cookie := os.Getenv("WPLACE_COOKIE")
	if cookie == "" {
		log.Fatal("WPLACE_COOKIE environment variable is required")
	}

	// Create a new client with your session cookie
	client := wplace.NewClient().WithCookie(cookie)

	// Paint pixels at coordinates (459,167) and (460,168) with color 8
	resp, err := client.PaintPixels(context.Background(), "turnstile-token", wplace.P(1222, 837), []wplace.Point{{X: 459, Y: 167}, {X: 460, Y: 168}}, []int{8, 8})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Success: %v\n", resp)
}
