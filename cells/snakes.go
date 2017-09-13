package cells

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/jessemillar/gautomata/tools"
)

func Snakes(m image.RGBA, w int, h int, palette []color.RGBA) {
	directions := []rune{'l', 'r', 'u', 'd'}
	magic := w * h / 5
	currentColor := palette[1]

	for i := 0; i < magic/3; i++ {
		// Start in a random spot
		x := rand.Intn(w)
		y := rand.Intn(h)
		// Pick a random color for each snake
		currentColor = tools.RandFromPalette(palette)
		// Pick a random starting direction
		currentDirection := directions[rand.Intn(len(directions))]

		// Make the initial dot
		m.Set(x, y, currentColor)

		for l := 0; l < magic; l++ {
			// Change directions every once in a while
			if rand.Intn(10) < 2 {
				currentDirection = directions[rand.Intn(len(directions))]
			}

			switch currentDirection {
			case 'l':
				x -= 1
				m.Set(x, y, currentColor)
			case 'r':
				x += 1
				m.Set(x, y, currentColor)
			case 'u':
				y -= 1
				m.Set(x, y, currentColor)
			case 'd':
				y += 1
				m.Set(x, y, currentColor)
			}
		}
	}
}
