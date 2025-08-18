package bot

import (
	"image"
	"image/color"

	"github.com/jrsap/wplace-worker/pkg/wplace"
)

const (
	PIXEL_IGNORED     = 0
	PIXEL_CORRECT     = 1
	PIXEL_INCORRECT   = 2
	PIXEL_IN_PROGRESS = 3
)

var (
	palette = []color.Color{
		wplace.ConvertColor("#000000"),
		wplace.ConvertColor("#006600"),
		wplace.ConvertColor("#aa0000"),
		wplace.ConvertColor("#aa33aa"),
	}
)

type imageStatus struct {
	tile                               wplace.Point
	pixel                              wplace.Point
	target, current                    image.Image
	diff                               *image.Paletted
	totalPixelCount, correctPixelCount int
}

func newImageStatus(target image.Image) *imageStatus {
	return &imageStatus{
		target: target,
		diff:   image.NewPaletted(target.Bounds(), palette),
	}
}

func (i *imageStatus) update(current image.Image) {
	i.current = current
	for y := 0; y < i.target.Bounds().Dy(); y++ {
		for x := 0; x < i.target.Bounds().Dx(); x++ {
			expected := i.target.At(x, y)
			if _, _, _, a := expected.RGBA(); a != 0 {
				i.diff.SetColorIndex(x, y, PIXEL_IGNORED)
				continue // Skip transparant pixels
			}
			i.totalPixelCount++
			if i.target.At(x, y) == i.current.At(x, y) {
				i.correctPixelCount++
				i.diff.SetColorIndex(x, y, PIXEL_CORRECT)
			}

			if i.diff.ColorIndexAt(x, y) != PIXEL_IN_PROGRESS {
				i.diff.SetColorIndex(x, y, PIXEL_INCORRECT)
			}
		}
	}
}

func (i *imageStatus) contains(tile, pixel wplace.Point) bool {
	// TODO
	return false
}
