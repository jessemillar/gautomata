package cells

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/jessemillar/gautomata/tools"
)

func Playground(m image.RGBA, w int, h int, palette []color.RGBA) {
	speechLength := w / 5
	background := palette[0]
	foreground := palette[1]

	// Create the initial image state where each pixel has an equal chance to be both states
	for x := 0; x < speechLength; x++ {
		for y := 0; y < h; y++ {
			if rand.Intn(3) == 1 {
				m.Set(x, y, foreground)
			}
		}
	}

	for y := 0; y < h; y++ {
		enthusiasm := 0
		applauseColor := tools.RandFromPalette(palette)

		for i := 0; i < speechLength; i++ {
			if m.At(i, y) == foreground {
				if rand.Intn(2) == 1 {
					enthusiasm++
				}
			}
		}

		for x := speechLength; x < w; x++ {
			if enthusiasm > 0 {
				leftNeighbor := m.At(x, y-1)
				rightNeighbor := m.At(x, y+1)
				interval := 2

				if enthusiasm < speechLength/2 {
					interval++
				}

				if rand.Intn(2) == 1 {
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
