package wplace

import "math"

func CalculateCoordinatesAtOffset(tile Point, pixel Point, offset Point) (Point, Point) {
	tileX := tile.X + (pixel.X+offset.X)/TileWidth
	tileY := tile.Y + (pixel.Y+offset.Y)/TileHeight

	pixelX := (pixel.X + offset.X) % TileWidth
	pixelY := (pixel.Y + offset.Y) % TileHeight

	return P(tileX, tileY), P(pixelX, pixelY)
}

func CalculateDistance(tile1, pixel1, tile2, pixel2 Point) float64 {
	// Convert tile coordinates to pixel coordinates
	x1 := tile1.X*TileWidth + pixel1.X
	y1 := tile1.Y*TileHeight + pixel1.Y
	x2 := tile2.X*TileWidth + pixel2.X
	y2 := tile2.Y*TileHeight + pixel2.Y

	// Calculate Euclidean distance
	dx := x2 - x1
	dy := y2 - y1
	distance := math.Sqrt(float64(dx*dx + dy*dy))

	return distance
}
