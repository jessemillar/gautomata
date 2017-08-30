package tools

import (
	"image/color"
	"math/rand"
)

func RandColor() color.Color {
	cap := 127

	red := uint32(rand.Intn(cap)) + 100
	green := uint32(rand.Intn(cap)) + 100
	blue := uint32(rand.Intn(cap)) + 100

	return color.RGBA{uint8(red), uint8(green), uint8(blue), 255}
}
