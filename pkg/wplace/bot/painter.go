package bot

import (
	"context"
	"math/rand"
	"net/http"
	"time"

	"github.com/avast/retry-go"
	"github.com/jrsap/wplace-worker/pkg/wplace"
)

func (b *Bot) painter(ctx context.Context, accountIdx int) {
	for {
		if ctx.Err() != nil {
			break
		}

		tile, pixels, colors := b.getNextPixels(accountIdx)
		b.log(accountIdx, "received %d pixels", len(pixels))

		if len(pixels) != 0 {
			if err := b.doPaint(ctx, accountIdx, tile, pixels, colors); err != nil {
				b.log(accountIdx, "error painting pixels: %v", err)
				b.cancelPixels(tile, pixels)
			} else {
				b.log(accountIdx, "painting succesful")
			}
			time.Sleep(time.Second * time.Duration(b.config.Limits.MinSecondsBetweenPaints))
		} else {
			b.log(accountIdx, "no work received")
		}
	}
}

func (b *Bot) doPaint(ctx context.Context, accountIdx int, tile wplace.Point, pixels []wplace.Point, colors []int) error {
	defer b.updateUserInfo(ctx, accountIdx)
	b.lock.Lock()
	defer b.lock.Unlock()

	var turnstileToken string
	var cookies []*http.Cookie
	err := retry.Do(
		func() error {
			var err error
			turnstileToken, cookies, err = b.cloudbuster.GetToken("https://wplace.live", "")
			return err
		},
		retry.DelayType(retry.BackOffDelay),
		retry.Delay(100*time.Millisecond),
		retry.Attempts(b.config.CloudBuster.MaxRetries))
	if err != nil {
		b.log(accountIdx, "failed to fetch CF token")
		return err
	}
	cookies = append(cookies, b.accounts[accountIdx].cookies...)

	b.wplaceClient.SetCookies(cookies)
	resp, err := b.wplaceClient.PaintPixels(ctx, turnstileToken, tile, pixels, colors)
	_ = resp
	return err
}

func (b *Bot) getNextPixels(accountIdx int) (tile wplace.Point, pixels []wplace.Point, colors []int) {
	pixelCount := b.config.Limits.MinPixelsPerRequest + rand.Intn(b.config.Limits.MaxPixelsPerRequest-b.config.Limits.MinPixelsPerRequest)
	pixelCount = min(b.accounts[accountIdx].userInfo.Charges.Max, pixelCount)
	b.log(accountIdx, "targeting %d pixels in next paint", pixelCount)

	charges := int(b.accounts[accountIdx].userInfo.Charges.Count)
	for charges < pixelCount {
		missingPixels := pixelCount - charges
		sleepTime := 30 * time.Second * time.Duration(missingPixels+rand.Intn(20))
		b.log(accountIdx, "not enough charges, sleeping %v", sleepTime)
		time.Sleep(sleepTime)
		charges = int(b.accounts[accountIdx].userInfo.Charges.Count)
	}

	for idx, img := range b.images {
		if b.config.Templates[idx].Disabled {
			continue
		}
		tile, pixels, colors = img.getWork(pixelCount)

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
