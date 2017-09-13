package cells

import (
	"image"
	"image/color"

	"github.com/jessemillar/gautomata/tools"
)

func Funnels(m image.RGBA, w int, h int, palette []color.RGBA) {
	foreground := palette[1]

	tools.RandTopLine(m, w, foreground)

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
