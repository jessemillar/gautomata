package tools

import (
	"image"
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

// RandState generates a random top line for the automata to feed off of
func RandState(m image.RGBA, w int, foreground color.Color) {
	for i := 0; i < w; i++ {
		upperBound := rand.Intn(10) + 10
		lowerBound := rand.Intn(1) + 1

		if rand.Intn(upperBound) < lowerBound {
			m.Set(i, 0, foreground)
		}
	}
}
