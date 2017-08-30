package tools

import (
	"image/color"
	"math/rand"
)

func RandColor() color.Color {
	cap := 127

	red := uint32(rand.Intn(cap)) + 127
	green := uint32(rand.Intn(cap)) + 127
	blue := uint32(rand.Intn(cap)) + 127

	return color.RGBA{uint8(red), uint8(green), uint8(blue), 255}
}

func GhostColor(inputColor color.Color) color.Color {
	fr, fg, fb, _ := inputColor.RGBA()
	return color.RGBA{uint8(fr), uint8(fg), uint8(fb), 127}
}
