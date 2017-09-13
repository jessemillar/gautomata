package cells

import (
	"image"
	"image/color"
	"math/rand"
)

func Flowers(m image.RGBA, w int, h int, palette []color.RGBA) {
	modifier := 15
	size := w * h * modifier
	background := palette[0]
	currentColor := palette[1]
	iterations := rand.Intn(size) + size

	// Pick a random pixel as the starting point
	m.Set(rand.Intn(w), rand.Intn(h), currentColor)

	for iterations > 0 {
		x := rand.Intn(w)
		y := rand.Intn(h)

		if m.At(x, y) == background {
			continue // Go to the next loop iteration
		}

		iterations--

		// Pick a random, non-background color every once in a while for variety
		if rand.Intn(size) < modifier*2 {
			currentColor = palette[rand.Intn(len(palette)-1)+1]
		}

		l := m.At(x-1, y) == background
		r := m.At(x+1, y) == background
		u := m.At(x, y-1) == background
		d := m.At(x, y+1) == background

		if l {
			m.Set(x-1, y, currentColor)
		}

		if r {
			m.Set(x+1, y, currentColor)
		}

		if u {
			m.Set(x, y-1, currentColor)
		}

		if d {
			m.Set(x, y+1, currentColor)
		}
	}
}
