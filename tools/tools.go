package tools

import (
	"image/color"
	"math/rand"
)

func RandColor() color.Color {
	red := uint32(rand.Intn(256))
	green := uint32(rand.Intn(256))
	blue := uint32(rand.Intn(256))

	return color.RGBA{uint8(red), uint8(green), uint8(blue), 255}
}

func LightenColor(inputColor color.Color, alpha int) color.Color {
	fr, fg, fb, _ := inputColor.RGBA()
	return color.RGBA{uint8(fr), uint8(fg), uint8(fb), uint8(alpha)}
}
