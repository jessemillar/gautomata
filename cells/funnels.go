package cells

import (
	"image"
	"image/color"
	"math/rand"
)

func Funnels(m image.RGBA, w int, h int, light color.Color, dark color.Color) {
	// Random initial state
	for i := 0; i < w; i++ {
		if rand.Intn(2) < 1 {
			m.Set(i, 0, light)
		}
	}

	// Loop through the y axis
	for y := 1; y < h; y++ {
		for x := 0; x < w; x++ {
			draw := 0

			if m.At(x-1, y-1) == dark {
				draw++
			}

			if m.At(x+1, y-1) == dark {
				draw++
			}

			if draw == 1 {
				m.Set(x, y, light)
			}
		}
	}
}
