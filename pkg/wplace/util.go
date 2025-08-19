package wplace

import (
	"image"
)

const TileHeight = 1000
const TileWidth = 1000

type Point image.Point

func P(x, y int) Point {
	return Point{
		X: x,
		Y: y,
	}
}

func CalculateCoordinatesAtOffset(tile Point, pixel Point, offset Point) (Point, Point) {
	tileX := tile.X + (pixel.X+offset.X)/TileWidth
	tileY := tile.Y + (pixel.Y+offset.Y)/TileHeight

	pixelX := (pixel.X + offset.X) % TileWidth
	pixelY := (pixel.Y + offset.Y) % TileHeight

	return P(tileX, tileY), P(pixelX, pixelY)
}

func CalculateOffset(tile1, pixel1, tile2, pixel2 Point) Point {
	// Convert tile coordinates to pixel coordinates
	x1 := tile1.X*TileWidth + pixel1.X
	y1 := tile1.Y*TileHeight + pixel1.Y
	x2 := tile2.X*TileWidth + pixel2.X
	y2 := tile2.Y*TileHeight + pixel2.Y
	dx := x2 - x1
	dy := y2 - y1

	return P(dx, dy)
}
