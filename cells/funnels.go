package cells

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/jessemillar/gautomata/tools"
)

func Funnels(m image.RGBA, w int, h int, background color.Color) {
	foreground := tools.GhostColor(background)

	// Random initial state
	for i := 0; i < w; i++ {
		if rand.Intn(2) < 1 {
			m.Set(i, 0, foreground)
		}
	}

	// Loop through the canvas
	for y := 1; y < h; y++ {
		for x := 0; x < w; x++ {
			draw := 0

			if m.At(x-1, y-1) == foreground {
				draw++
			}

			if m.At(x+1, y-1) == foreground {
				draw++
			}

			if draw == 1 {
				m.Set(x, y, foreground)
			}
		}
	}
}
