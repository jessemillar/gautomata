package cells

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/jessemillar/gautomata/tools"
)

func Playground(m image.RGBA, w int, h int, palette []color.RGBA) {
	speechLength := w / 5
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
		impressions := 0
		applauseColor := tools.RandFromPalette(palette)

		for i := 0; i < speechLength; i++ {
			if m.At(i, y) == foreground {
				impressions++
			}
		}

		enthusiasm := rand.Intn(w) / impressions

		for x := speechLength; x < w; x++ {
			if enthusiasm > 0 {
				m.Set(x, y, applauseColor)

				if rand.Intn(5) == 1 {
					enthusiasm--
				}
			}
		}
	}
}
