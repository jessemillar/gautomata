package tools

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/lucasb-eyer/go-colorful"
)

func RandPalette(count int) ([]color.RGBA, error) {
	generatedPalette, err := colorful.SoftPalette(count)
	if err != nil {
		return []color.RGBA{}, err
	}

	finalPalette := []color.RGBA{}

	for _, c := range generatedPalette {
		finalPalette = append(finalPalette, color.RGBA{uint8(c.R * 255), uint8(c.G * 255), uint8(c.B * 255), 255})
	}

	return finalPalette, nil
}

// RandTopLine generates a random top line for the automata to feed off of
func RandTopLine(m image.RGBA, w int, foreground color.RGBA) {
	for i := 0; i < w; i++ {
		upperBound := rand.Intn(10) + 10
		lowerBound := rand.Intn(1) + 1

		if rand.Intn(upperBound) < lowerBound {
			m.Set(i, 0, foreground)
		}
	}
}
