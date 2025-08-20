package bot

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/avast/retry-go"
)

func (b *Bot) withCFClearance(ctx context.Context, otherCookies []*http.Cookie, force bool) ([]*http.Cookie, error) {
	cookies := []*http.Cookie{}
	var headers http.Header
	var err error

	if !force {
		for _, c := range otherCookies {
			if c.Name == "cf_clearance" {
				if c.Expires.After(time.Now().Add(-time.Minute)) {
					return otherCookies, nil
				}
			}
		}
	}

	log.Printf("refreshing cloudflare cookie")

	if b.config.Byparr != nil {
		err = retry.Do(func() error {
			var err error
			cookies, headers, err = b.byparrClient.GetAuthentication(ctx, "https://wplace.live", time.Duration(*b.config.Byparr.MaxDuration))
			return err
		}, retry.DelayType(retry.BackOffDelay), retry.Delay(100*time.Millisecond), retry.Attempts(b.getCfRetries()))
	}

	newNames := map[string]bool{}
	for _, c := range cookies {
		newNames[c.Name] = true
	}

	for _, c := range otherCookies {
		if !newNames[c.Name] {
			cookies = append(cookies, c)
		}
	}

	if ua, ok := headers["User-Agent"]; ok {
		b.wplaceClient.WithUserAgent(strings.Join(ua, "; "))
	}

	log.Printf("refreshing cloudflare cookie done!")

	return cookies, err
}

func (b *Bot) getCfRetries() uint {
	if b.config.Byparr != nil && b.config.Byparr.MaxRetries == nil {
		return 3
	}

	return uint(*b.config.Byparr.MaxRetries)
}
