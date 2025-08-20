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
	cookies, err := http.ParseCookie("cf_clearance=bCR7FQ5dNuFPsQD5BdT1aOj6Ub7hjNThXp0_mFimu3Q-1755685216-1.2.1.1-vw306MRXVZa_Hd0XrqIdeH9Tz5i7twhbYAQsMIl94solUs69mBqMDPAzTjOMTovNdQtmuP2VLbH8MdlYbiNOmg1FKHUEtJ5GBSjfxgiQu_cj_981IxwgIYkBN9ZJBXKqZFulTt2BJcl88JMwF2y1WyPRRIpTNVHsXEXBCVq0BWRH_WluZgbJzMHY.f0hw9TZYpVEq9Be7U2UiCyHdSsJn1.fWBz09ieyJNR5ql56r_I; j=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjg1ODU0NDgsInNlc3Npb25JZCI6IlU1ZFllWmxOTGtrUDM3bEtrb3dpaFpua01nZlVDOHl6cnZzVjVhMExEM2c9IiwiaXNzIjoid3BsYWNlIiwiZXhwIjoxNzU4MTk3NzczLCJpYXQiOjE3NTU2MDU3NzN9.sHoJkU49t8p1tXylE2q7YIToTIhUHI-cL6kNfdJDRWs")
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
