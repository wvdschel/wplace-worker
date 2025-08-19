package bot

import (
	"image"
	"image/color"
	"sync"

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

	lock *sync.Mutex
}

func newImageStatus(tile, pixel wplace.Point, target image.Image) *imageStatus {
	return &imageStatus{
		target: target,
		diff:   image.NewPaletted(target.Bounds(), palette),
		lock:   &sync.Mutex{},
		tile:   tile,
		pixel:  pixel,
	}
}

func (i *imageStatus) getImage() image.Image {
	res := image.NewRGBA(image.Rect(0, 0, i.target.Bounds().Dx()*3, i.target.Bounds().Dy()))

	wplace.BlitImage(i.target, res, i.target.Bounds(), image.Point{
		X: 0,
		Y: 0,
	})
	wplace.BlitImage(i.current, res, i.target.Bounds(), image.Point{
		X: i.target.Bounds().Dx(),
		Y: 0,
	})
	wplace.BlitImage(i.diff, res, i.target.Bounds(), image.Point{
		X: i.target.Bounds().Dx() * 2,
		Y: 0,
	})

	return res
}

func (i *imageStatus) update(current image.Image) {
	i.lock.Lock()
	defer i.lock.Unlock()

	i.current = current
	for y := 0; y < i.target.Bounds().Dy(); y++ {
		for x := 0; x < i.target.Bounds().Dx(); x++ {
			expected := i.target.At(x, y)
			if _, _, _, a := expected.RGBA(); a == 0 {
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
	offset := wplace.CalculateOffset(i.tile, i.pixel, tile, pixel)

	return offset.X < i.target.Bounds().Dx() && offset.Y < i.target.Bounds().Dy()
}

func (i *imageStatus) getWork(maxCount int) (wplace.Point, []wplace.Point) {
	i.lock.Lock()
	defer i.lock.Unlock()

	tile := wplace.P(-1, -1)
	pixels := make([]wplace.Point, 0, maxCount)
	for y := i.diff.Bounds().Min.Y; y < i.diff.Bounds().Max.Y; y++ {
		if len(pixels) >= maxCount {
			break
		}
		for x := i.diff.Bounds().Min.X; x < i.diff.Bounds().Max.X; x++ {
			if len(pixels) >= maxCount {
				break
			}
			if i.diff.ColorIndexAt(x, y) != PIXEL_INCORRECT {
				continue
			}
			p := wplace.P(x, y)
			p_tile, p := wplace.CalculateCoordinatesAtOffset(i.tile, i.pixel, p)

			if tile.X == -1 {
				tile = p_tile
			}
			if tile != p_tile {
				continue
			}

			i.diff.SetColorIndex(x, y, PIXEL_IN_PROGRESS)
			pixels = append(pixels, p)
		}
	}

	return tile, pixels
}

func (i *imageStatus) returnWork(tile wplace.Point, pixels []wplace.Point) {
	i.lock.Lock()
	defer i.lock.Unlock()

	for _, p := range pixels {
		pointInImage := wplace.CalculateOffset(i.tile, i.pixel, tile, p)

		i.diff.SetColorIndex(pointInImage.X, pointInImage.Y, PIXEL_INCORRECT)
	}
}
