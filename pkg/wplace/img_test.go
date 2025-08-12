package wplace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Colors(t *testing.T) {
	assert.Len(t, ColorPallet, 31)
}
