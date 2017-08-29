package tools

import (
	"image/color"
	"math/rand"
)

func RandColor() color.Color {
	cap := 150

	red := uint32(rand.Intn(cap))
	green := uint32(rand.Intn(cap))
	blue := uint32(rand.Intn(cap))

	return color.RGBA{uint8(red), uint8(green), uint8(blue), 255}
}

func LightenColor(inputColor color.Color) color.Color {
	fr, fg, fb, _ := inputColor.RGBA()
	return color.RGBA{uint8(fr), uint8(fg), uint8(fb), 200}
}
