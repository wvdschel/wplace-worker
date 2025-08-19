package bot

import (
	"context"
	"net/http"
	"time"

	"github.com/avast/retry-go"
)

func (b *Bot) withCFClearance(ctx context.Context, otherCookies []*http.Cookie) ([]*http.Cookie, error) {
	cookies := []*http.Cookie{}
	var err error

	if b.config.Byparr != nil {
		err = retry.Do(func() error {
			var err error
			cookies, err = b.byparrClient.GetCookie(ctx, "https://wplace.live", time.Duration(*b.config.Byparr.MaxDuration))
			return err
		}, retry.DelayType(retry.BackOffDelay), retry.Delay(100*time.Millisecond), retry.Attempts(b.getCfRetries()))
	}

	cookies = append(cookies, otherCookies...)

	return cookies, err
}

func (b *Bot) getCfRetries() uint {
	if b.config.Byparr != nil && b.config.Byparr.MaxRetries == nil {
		return 3
	}

	return uint(*b.config.Byparr.MaxRetries)
}
