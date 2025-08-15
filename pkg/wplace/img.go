package wplace

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"os"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

var ColorPallette []color.Color = []color.Color{
	convertColor("#000000"),
	convertColor("#3c3c3c"),
	convertColor("#787878"),
	convertColor("#d2d2d2"),
	convertColor("#ffffff"),
	convertColor("#600018"),
	convertColor("#ed1c24"),
	convertColor("#ff7f27"),
	convertColor("#f6aa09"),
	convertColor("#f9dd3b"),
	convertColor("#fffabc"),
	convertColor("#0eb968"),
	convertColor("#13e67b"),
	convertColor("#87ff5e"),
	convertColor("#0c816e"),
	convertColor("#13e1be"),
	convertColor("#13e1be"),
	convertColor("#28509e"),
	convertColor("#4093e4"),
	convertColor("#60f7f2"),
	convertColor("#6b50f6"),
	convertColor("#99b1fb"),
	convertColor("#780c99"),
	convertColor("#aa38b9"),
	convertColor("#e09ff9"),
	convertColor("#cb007a"),
	convertColor("#ec1f80"),
	convertColor("#f38da9"),
	convertColor("#684634"),
	convertColor("#95682a"),
	convertColor("#f8b277"),
	convertColor("#00000000"),
}

func convertColor(hex string) color.Color {
	// Takes a hex string in the form of "#rrggbbaa", "#rrggbb", "rrggbb", "rrggbbaa" and returns a color
	var r, g, b, a uint8
	a = 255 // default alpha to opaque

	if hex[0] == '#' {
		hex = hex[1:]
	}

	switch len(hex) {
	case 6:
		// #rrggbb
		fmt.Sscanf(hex, "%02x%02x%02x", &r, &g, &b)
	case 8:
		// #rrggbbaa
		fmt.Sscanf(hex, "%02x%02x%02x%02x", &r, &g, &b, &a)
	default:
		// invalid format, return black
		return color.RGBA{0, 0, 0, 255}
	}

	return color.RGBA{r, g, b, a}
}

func LoadImage(path string) (image.Image, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	return img, err
}

func ConvertToPallette(img image.Image) image.PalettedImage {
	bounds := img.Bounds()
	palImg := image.NewPaletted(bounds, ColorPallette)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			target := img.At(x, y)

			minDiff := math.MaxFloat64
			bestMatchIdx := 4 // White pixel
			//bestMatchIdx := 31 // Transparant pixel
			for idx, col := range ColorPallette {
				diff := colorDistance(target, col)
				if diff < minDiff {
					minDiff = diff
					bestMatchIdx = idx
				}
			}
			//fmt.Printf("best match for %v is %v with diff %v\n", target, ColorPallette[bestMatchIdx], minDiff)
			// if minDiff > 0.1 {
			// 	fmt.Printf("Approximating %v to %v (%v difference)\n", target, ColorPallette[bestMatchIdx], minDiff)
			// }
			palImg.SetColorIndex(x, y, uint8(bestMatchIdx))
		}
	}
	return palImg
}

func colorDistance(c1, c2 color.Color) float64 {
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()
	res := math.Sqrt(float64((r1-r2)*(r1-r2) + (g1-g2)*(g1-g2) + (b1-b2)*(b1-b2) + (a1-a2)*(a1-a2)))
	//fmt.Printf("color diff between %v and %v: %v\n", c1, c2, res)
	return res
}

func ScaleImage(img image.Image, factor float64) image.Image {
	// Scales an image up or down by factor.
	// When upscaling, the pixels will be interpolated linearly from their nearest neighbours in the source image.
	// When downscaling, the pixels will be a weighted average of the source pixels they contain, partially or in full.
	if factor <= 0 {
		return image.NewRGBA(image.Rect(0, 0, 0, 0))
	}

	srcBounds := img.Bounds()
	srcW := srcBounds.Dx()
	srcH := srcBounds.Dy()

	newW := int(float64(srcW) * factor)
	newH := int(float64(srcH) * factor)

	if newW <= 0 || newH <= 0 {
		return image.NewRGBA(image.Rect(0, 0, 0, 0))
	}

	newBounds := image.Rect(0, 0, newW, newH)
	newImg := image.NewRGBA(newBounds)

	if factor < 1 {
		// Downscaling: use weighted average of source pixels
		for i := 0; i < newW; i++ {
			for j := 0; j < newH; j++ {
				// Calculate source rectangle for current destination pixel
				sx0 := float64(i) / factor
				sy0 := float64(j) / factor
				sx1 := sx0 + 1.0/factor
				sy1 := sy0 + 1.0/factor

				// Get integer bounds for source pixels
				ix0 := int(sx0)
				iy0 := int(sy0)
				ix1 := int(sx1)
				iy1 := int(sy1)

				var totalR, totalG, totalB, totalA uint32
				count := 0

				for ix := ix0; ix < ix1; ix++ {
					for iy := iy0; iy < iy1; iy++ {
						if ix >= srcBounds.Min.X && ix < srcBounds.Max.X &&
							iy >= srcBounds.Min.Y && iy < srcBounds.Max.Y {
							c := img.At(ix, iy)
							r, g, b, a := c.RGBA()
							totalR += r
							totalG += g
							totalB += b
							totalA += a
							count++
						}
					}
				}

				if count == 0 {
					// No pixels found, use transparent
					newImg.SetRGBA(i, j, color.RGBA{0, 0, 0, 0})
					continue
				}

				// Average the components
				avgR := totalR / uint32(count)
				avgG := totalG / uint32(count)
				avgB := totalB / uint32(count)
				avgA := totalA / uint32(count)

				newImg.SetRGBA(i, j, color.RGBA{
					R: uint8(avgR >> 8),
					G: uint8(avgG >> 8),
					B: uint8(avgB >> 8),
					A: uint8(avgA >> 8),
				})
			}
		}
	} else {
		// Upscaling: use nearest neighbor
		for i := 0; i < newW; i++ {
			for j := 0; j < newH; j++ {
				// Map destination coordinates to source
				x := srcBounds.Min.X + int(float64(i)/factor)
				y := srcBounds.Min.Y + int(float64(j)/factor)

				// Clamp to source bounds
				if x < srcBounds.Min.X {
					x = srcBounds.Min.X
				} else if x >= srcBounds.Max.X {
					x = srcBounds.Max.X - 1
				}

				if y < srcBounds.Min.Y {
					y = srcBounds.Min.Y
				} else if y >= srcBounds.Max.Y {
					y = srcBounds.Max.Y - 1
				}

				c := img.At(x, y)
				r, g, b, a := c.RGBA()
				newImg.SetRGBA(i, j, color.RGBA{
					R: uint8(r >> 8),
					G: uint8(g >> 8),
					B: uint8(b >> 8),
					A: uint8(a >> 8),
				})
			}
		}
	}

	return newImg
}

func blitImage(src image.Image, dest *image.RGBA, srcBounds image.Rectangle, destPos image.Point) {
	// Copy part of src contained within srcBounds into the part of dest contained by destBounds
	for y := 0; y < srcBounds.Dy(); y++ {
		srcY := srcBounds.Min.Y + y
		dstY := destPos.Y + y
		if dstY >= dest.Bounds().Max.Y || srcY >= src.Bounds().Max.Y {
			continue
		}
		for x := 0; x < srcBounds.Dx(); x++ {
			srcX := srcBounds.Min.X + x
			dstX := destPos.X + x
			if dstX >= dest.Bounds().Max.X || srcX >= src.Bounds().Max.X {
				continue
			}
			dest.Set(dstX, dstY, src.At(srcX, srcY))
			if x == 0 || y == 0 {
				dest.Set(dstX, dstY, color.White)
			}
		}
	}
}
