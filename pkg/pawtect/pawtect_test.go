package pawtect_test

import (
	"testing"

	"github.com/wvdschel/wplace-worker/pkg/pawtect"

	"github.com/stretchr/testify/require"
)

func TestSetUserID(t *testing.T) {
	p, err := pawtect.Load()
	require.NoError(t, err)

	p.SetUserId(12345)
}
