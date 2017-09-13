package cells

import (
	"image"
	"image/color"

	"github.com/jessemillar/gautomata/tools"
)

func Rule110(m image.RGBA, w int, h int, palette []color.RGBA) {
	background := palette[0]
	foreground := palette[1]

	tools.RandTopLine(m, w, foreground)

	// Loop through the canvas
	for y := 1; y < h; y++ {
		for x := 0; x < w; x++ {
			l := m.At(x-1, y-1) == foreground
			c := m.At(x, y-1) == foreground
			r := m.At(x+1, y-1) == foreground

			if l && c && r {
				m.Set(x, y, background)
				continue
			}

			if l && c && !r {
				m.Set(x, y, foreground)
				continue
			}

			if l && !c && r {
				m.Set(x, y, foreground)
				continue
			}

			if l && !c && !r {
				m.Set(x, y, background)
				continue
			}

			if !l && c && r {
				m.Set(x, y, foreground)
				continue
			}

			if !l && c && !r {
				m.Set(x, y, foreground)
				continue
			}

			if !l && !c && r {
				m.Set(x, y, foreground)
				continue
			}

			if !l && !c && !r {
				m.Set(x, y, background)
			}
		}
	}
}
