package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	// Rect is the image's bounds.
	rect image.Rectangle
}

func (im Image) Bounds() image.Rectangle {
	return im.rect
}

func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (im Image) At(x, y int) color.Color {
	r := uint8((x + y) / 2)
	g := uint8(x * y)
	b := uint8(x ^ y)
	return color.RGBA{r, g, b, 255}
}

func main() {
	m := Image{
		rect: image.Rectangle{
			Min: image.Point{0, 0},
			Max: image.Point{100, 100},
		},
	}
	pic.ShowImage(m)
}
