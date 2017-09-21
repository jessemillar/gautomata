package cells

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/jessemillar/gautomata/tools"
)

func Applause(m image.RGBA, w int, h int, palette []color.RGBA) {
	speechLength := h / 5
	background := palette[0]
	foreground := palette[1]

	// Create the initial image state where each pixel in a band has a chance to be both states
	for y := 0; y < speechLength; y++ {
		for x := 0; x < w; x++ {
			if rand.Intn(3) == 1 {
				m.Set(x, y, foreground)
			}
		}
	}

	for x := 0; x < w; x++ {
		enthusiasm := 0
		applauseColor := tools.RandFromPalette(palette)

		for i := 0; i < speechLength; i++ {
			if m.At(x, i) == foreground {
				if rand.Intn(2) == 1 {
					if m.At(x, i-1) == foreground {
						enthusiasm += 2
					} else {
						enthusiasm++
					}
				}
			}
		}

		for y := speechLength; y < h; y++ {
			if enthusiasm > 0 {
				leftNeighbor := m.At(x-1, y)
				rightNeighbor := m.At(x+1, y)
				interval := 2

				if enthusiasm < speechLength/10 {
					interval++
				}

				if rand.Intn(interval) == 1 {
					m.Set(x, y, applauseColor)
				}

				if rand.Intn(5) == 1 {
					if leftNeighbor != background && rightNeighbor != background {
						if rand.Intn(2) == 1 {
							enthusiasm++
						}
					} else if leftNeighbor != background || rightNeighbor != background {
						if rand.Intn(4) == 1 {
							enthusiasm++
						}
					} else {
						enthusiasm--
					}
				}
			}
		}
	}
}
