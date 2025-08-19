package bot

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jrsap/wplace-worker/pkg/wplace"
	"github.com/stretchr/testify/assert"
)

func TestComparePoints(t *testing.T) {
	p1 := wplace.P(101, 102)
	p2 := wplace.P(101, 102)

	assert.True(t, p1 == p2)
}

func TestPrintCookie(t *testing.T) {
	cookies, err := http.ParseCookie("cf_clearance=AlHczDB7u6mqQS7BPoR4T5lmNZbpJQSX.u_9T68C9p0-1755588223-1.2.1.1-RMpfwqYz30GSrSjMtaxRmhb0mq4xBoYoeoAYmSROywbN_dzd1r9XCYkOvNsxDvshLUdh4qzm6RQEebeMNyEIJUCeWDez9dpeRmvHq.YNaiYRJiGGbYn_urbLi6K7STh5tXXt2aYAP9WWXoTMXNEIpV9hJ3oRCmZfZMgabn.SyZL2MZR2F36sJHM7mNBmHRYvqflU9R.kjtsSM6.P50rLlwTMPhkKpf8true3mHGTisI; j=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjg1ODU0NDgsInNlc3Npb25JZCI6IlU1ZFllWmxOTGtrUDM3bEtrb3dpaFpua01nZlVDOHl6cnZzVjVhMExEM2c9IiwiaXNzIjoid3BsYWNlIiwiZXhwIjoxNzU4MTk3NzczLCJpYXQiOjE3NTU2MDU3NzN9.sHoJkU49t8p1tXylE2q7YIToTIhUHI-cL6kNfdJDRWs")
	assert.NoError(t, err)

	cleanCookies := ""
	for _, c := range cookies {
		if c.Name != "cf_clearance" {
			cleanCookies = fmt.Sprintf("%s; ", c)
		}
	}
	cleanCookies = cleanCookies[0 : len(cleanCookies)-2]
	fmt.Println(cleanCookies)
}
