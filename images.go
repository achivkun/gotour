package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		result[i] = make([]uint8, dx)
	}

	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			result[x][y] = uint8(x ^ y)
		}
	}
	return result
}

type Image struct {
	w, h int
}

func (im Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, im.w, im.h)
}

func (im Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (im Image) At(x, y int) color.Color {
	v := uint8(x ^ y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{128, 128}
	pic.ShowImage(m)
}
