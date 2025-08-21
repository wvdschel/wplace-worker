package bot

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/jrsap/wplace-worker/pkg/wplace"
)

const turnstileToken = "0.-IvGWVcng4C1p8YUqNoteDWSX0FHv3dmMAqstTjd0AhOBddIaJi_mKBaWSOEMnqvpfvmYyq3Ij24pTPbQ0-g6pUKUdyY_vf0OcEyNYIPk7OVpMCUrjkg9kf9bT_M10XGEa59ndh-NnLhchApfd9gaAwDHh536xmJNY9wtD1nzMgpxLBAvghUsAnqCG8ngcMjRyBSUI1WTHrEDtZ4NUuLKpGV4jWVY3q2-iBsQnsLenbx7FfuzDGmNiVLEs636IKxpGtIi-Q0KBnh4j5oYGRlmF3lfNGIPJpxtk2grvHUyUDIT3V2FT1pqRd5MGucasc_NPatWl42t4z4FQhPUMFEXxEHRDF7UG0eEDErZZ-nv3AUNSL4mFxHII0HkWNEhfbyL7IgrB2uqvMqmalqeUKGk3ZUTjgbnp22gK6p1P81tyPw1-AmPKTOwnHDcpw1wMpBXB-Ljq7GJtV0-ZZatUfGW5DQ6zluwpDcq0MPkaHHv9AV8RqUcTUgppSrqwR47QWhlESJY25Nr5bFEAfPspS2X6EBxXiEkRELm9uI4qGPQ-tSU2QQIlTwH-aE_tfH9pvWnYNSe2BEVJfnMnh2r2NhChn9sqS3n1Yss0sSZX85Rm4mgZix0GXpOFuSHYt0J5IjWl3hQ-DXvoiclyrfcTjGZx8XMHkxpGpS2JWR6Hvtq6DWqRfQkzFBZQj57quQEgT_EWS4-kMpa7h2LZzaKn9Ndvva5iIC6YXQ3p7r6OIPfdRYq9ebjtm3yA1qfI2tmDIQ_u9ktWrXMhPIqxwT7Rcw1VIBXjSwlrmc187JtpMnFyN01EBldP18-RL5f-9SMn6__wBgKSvVD6TpIe-OjOlTSMANSFTA4G9QCi0j1sqRhrE.eNOIEGGiTzet6TPpM9aC1Q.1c839e9561177c3b9ddc62fb1d611f689451ce9f27d72ddad27f5130ce349175"

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

	resp, err := b.wplaceClient.PaintPixels(ctx, turnstileToken, tile, pixels, colors)
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
