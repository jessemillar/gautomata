package cells

import (
	"image"
	"image/color"
	"math/rand"
)

func Rule110(m image.RGBA, w int, h int, light color.Color, dark color.Color) {
	// Random initial state
	for i := 0; i < w; i++ {
		if rand.Intn(2) < 1 {
			m.Set(i, 0, light)
		}
	}

	// Loop through the canvas
	for y := 1; y < h; y++ {
		for x := 0; x < w; x++ {
			left := m.At(x-1, y-1) == light
			center := m.At(x, y-1) == light
			right := m.At(x+1, y-1) == light

			if left && center && !right {
				m.Set(x, y, light)
				continue
			}

			if left && !center && right {
				m.Set(x, y, light)
				continue
			}

			if !left && center && right {
				m.Set(x, y, light)
				continue
			}

			if !left && center && !right {
				m.Set(x, y, light)
				continue
			}

			if !left && !center && right {
				m.Set(x, y, light)
			}
		}
	}
}
