package cells

import (
	"image"
	"image/color"
	"math/rand"
)

func Playground(m image.RGBA, w int, h int, palette []color.RGBA) {
	// http://www.emanueleferonato.com/2011/05/17/using-cellular-automata-to-generate-random-land-and-water-maps-with-flash/

	background := palette[0]
	foreground := palette[1]

	// Create the initial image state where each pixel has an equal chance to be both states
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if rand.Intn(2) == 1 {
				m.Set(x, y, foreground)
			}
		}
	}

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			waterCount := 0
			isWater := m.At(x, y) == background

			if m.At(x-1, y) == background {
				waterCount++
			}
			if m.At(x+1, y) == background {
				waterCount++
			}
			if m.At(x, y-1) == background {
				waterCount++
			}
			if m.At(x, y+1) == background {
				waterCount++
			}

			if m.At(x-1, y-1) == background {
				waterCount++
			}
			if m.At(x+1, y-1) == background {
				waterCount++
			}
			if m.At(x-1, y+1) == background {
				waterCount++
			}
			if m.At(x+1, y+1) == background {
				waterCount++
			}

			if isWater {
				if waterCount <= 3 {
					m.Set(x, y, foreground)
				}
			} else {
				if waterCount >= 5 {
					m.Set(x, y, background)
				}
			}
		}
	}
}
