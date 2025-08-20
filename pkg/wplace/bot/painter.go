package bot

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/jrsap/wplace-worker/pkg/wplace"
)

func (b *Bot) painter(ctx context.Context, cookieIndex int) {
	for {
		if ctx.Err() != nil {
			break
		}

		tile, pixels, colors := b.getNextPixels(cookieIndex)
		log.Printf("painter %d: received %d pixels", cookieIndex, len(pixels))

		if len(pixels) != 0 {
			if err := b.refreshCloudFlareToken(ctx, cookieIndex, false); err != nil {
				log.Printf("painter %d: error refreshing CloudFlare token: %v\n", cookieIndex+1, err)
			}

			if err := b.doPaint(ctx, cookieIndex, tile, pixels, colors); err != nil {
				log.Printf("painter %d: error painting pixels: %v\n", cookieIndex+1, err)
				b.cancelPixels(tile, pixels)
			} else {
				log.Printf("painter %d: painting succesfull", cookieIndex)
			}
		} else {
			log.Printf("painter %d: no work received", cookieIndex)
		}

		timeDiff := b.config.Limits.MaxSecondsBetweenPaints - b.config.Limits.MinSecondsBetweenPaints
		sleepSeconds := rand.Intn(timeDiff) + b.config.Limits.MinSecondsBetweenPaints
		sleepTime := time.Duration(sleepSeconds) * time.Second
		log.Printf("painter %d: sleeping for %v", cookieIndex, sleepTime)
		time.Sleep(sleepTime)
	}
}

func (b *Bot) refreshCloudFlareToken(ctx context.Context, cookieIdx int, force bool) error {
	b.lock.Lock()
	defer b.lock.Unlock()

	cookie, err := b.withCFClearance(ctx, b.cookies[cookieIdx], force)
	if err != nil {
		return err
	}
	b.cookies[cookieIdx] = cookie
	return nil
}

func (b *Bot) doPaint(ctx context.Context, cookieIdx int, tile wplace.Point, pixels []wplace.Point, colors []int) error {
	b.lock.Lock()
	defer b.lock.Unlock()

	b.wplaceClient.SetCookies(b.cookies[cookieIdx])

	resp, err := b.wplaceClient.PaintPixels(ctx, tile, pixels, colors)
	// if resp != nil && resp.Error == "refresh" {
	// 	b.lock.Unlock()
	// 	b.refreshCloudFlareToken(ctx, cookieIdx, true)
	// 	b.lock.Lock()
	// }
	_ = resp
	return err
}

func (b *Bot) getWorkCount(cookieIdx int) int {
	b.lock.Lock()
	defer b.lock.Unlock()

	charges := int(b.accounts[cookieIdx].Charges.Count)
	if charges < b.config.Limits.MaxPixelsPerRequest {
		return charges
	}

	return b.config.Limits.MaxPixelsPerRequest
}

func (b *Bot) getNextPixels(cookieIdx int) (tile wplace.Point, pixels []wplace.Point, colors []int) {
	for _, img := range b.images {

		tile, pixels, colors = img.getWork(b.getWorkCount(cookieIdx))

		if len(pixels) != 0 {
			return tile, pixels, colors
		}
	}

	return wplace.P(0, 0), nil, nil
}

func (b *Bot) cancelPixels(tile wplace.Point, pixels []wplace.Point) {
	if len(pixels) == 0 {
		return
	}

	for _, img := range b.images {
		if img.contains(tile, pixels[0]) {
			img.returnWork(tile, pixels)
			return
		}
	}
}
