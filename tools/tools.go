package tools

import (
	"image/color"
	"math/rand"
)

func RandColor(mix color.Color) color.Color {
	red := uint32(rand.Intn(256))
	green := uint32(rand.Intn(256))
	blue := uint32(rand.Intn(256))
	mixR, mixG, mixB, _ := mix.RGBA()

	if mix != nil {
		red = (red + mixR) / 2
		green = (green + mixG) / 2
		blue = (blue + mixB) / 2
	}

	return color.RGBA{uint8(red), uint8(green), uint8(blue), 255}
}
