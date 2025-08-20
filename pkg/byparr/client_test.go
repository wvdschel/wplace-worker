package byparr

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWplaceCloudflareCookies(t *testing.T) {
	c, err := New()
	assert.NoError(t, err)

	cookies, _, err := c.GetAuthentication(context.Background(), "https://wplace.live", 30*time.Second)
	assert.NoError(t, err)

	cfClearanceFound := false
	for _, cookie := range cookies {
		fmt.Println(cookie.Name, cookie.Value)
		if cookie.Name == "cf_clearance" {
			cfClearanceFound = true
		}
	}
	assert.True(t, cfClearanceFound)
}
