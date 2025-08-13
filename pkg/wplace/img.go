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
