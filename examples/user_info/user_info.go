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
	client := wplace.NewClient().WithCookie(cookie) //.WithBaseURL("https://localhost:8443")

	// Paint pixels at coordinates (459,167) and (460,168) with color 8
	resp, err := client.FetchUserInfo(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s: %0.0f/%d pixels available\n", resp.Name, resp.Charges.Count, resp.Charges.Max)
}
