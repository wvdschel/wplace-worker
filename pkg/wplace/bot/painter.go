package bot

import (
	"context"
	"encoding/json"
	"fmt"
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
		b.log(accountIdx, "painting %d pixels", len(pixels))

		if len(pixels) != 0 {
			if resp, err := b.doPaint(ctx, accountIdx, tile, pixels, colors); err != nil {
				b.log(accountIdx, "error painting pixels: %v", err)
				b.cancelPixels(tile, pixels)
			} else {
				err := retry.Do(
					func() error {
						b.refreshImages(ctx)
						correctPixels := 0

						for _, img := range b.images {
							if !img.contains(tile, pixels[0]) {
								continue
							}

							for _, px := range pixels {
								if img.getPixelStatus(tile, px) == PIXEL_CORRECT {
									correctPixels += 1
								}
							}
							break
						}

						if correctPixels != len(pixels) {
							return fmt.Errorf("only %d / %d pixels updated correctly", correctPixels, len(pixels))
						}
						return nil
					},
					retry.Attempts(5),
					retry.DelayType(retry.BackOffDelay),
					retry.Delay(1*time.Second))
				if err != nil {
					data, _ := json.Marshal(resp)
					b.log(accountIdx, "failed to paint pixels: %s", err.Error())
					b.log(accountIdx, "response: %s", string(data))
				} else {
					b.log(accountIdx, "painting succesful")
				}
			}
		} else {
			b.log(accountIdx, "no work received")
		}
		time.Sleep(time.Second * time.Duration(b.config.Limits.MinSecondsBetweenPaints))
	}
}

func (b *Bot) doPaint(ctx context.Context, accountIdx int, tile wplace.Point, pixels []wplace.Point, colors []int) (*wplace.PixelResponse, error) {
	defer b.updateUserInfo(ctx, accountIdx)
	b.accounts[accountIdx].lock.Lock()
	defer b.accounts[accountIdx].lock.Unlock()

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
		return nil, err
	}
	cookies = append(cookies, b.accounts[accountIdx].cookies...)

	b.accounts[accountIdx].client.SetCookies(cookies)
	resp, err := b.accounts[accountIdx].client.PaintPixels(ctx, turnstileToken, tile, pixels, colors)
	return resp, err
}

func (b *Bot) getNextPixels(accountIdx int) (tile wplace.Point, pixels []wplace.Point, colors []int) {
	pixelCount := b.config.Limits.MinPixelsPerRequest + rand.Intn(b.config.Limits.MaxPixelsPerRequest-b.config.Limits.MinPixelsPerRequest)
	pixelCount = min(b.accounts[accountIdx].userInfo.Charges.Max, pixelCount)
	b.log(accountIdx, "targeting %d pixels in next paint", pixelCount)

	charges := int(b.accounts[accountIdx].userInfo.Charges.Count)
	for charges < pixelCount {
		missingPixels := pixelCount - charges
		sleepTime := 30 * time.Second * time.Duration(missingPixels)
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
