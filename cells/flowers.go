package cells

import (
	"image"
	"image/color"
	"math/rand"
)

func Flowers(m image.RGBA, w int, h int, palette []color.RGBA) {
	currentColor := palette[1]

	// Pick a random pixel as the starting point
	m.Set(rand.Intn(w), rand.Intn(h), currentColor)

	for !canvasIsFull(m, w, h, palette[0]) {
		x := rand.Intn(w)
		y := rand.Intn(h)

		if m.At(x, y) == palette[0] {
			continue // Go to the next loop iteration
		}

		if rand.Intn(w*h/2) < 5 {
			currentColor = palette[rand.Intn(len(palette)-1)+1]
		}

		l := m.At(x-1, y) == palette[0]
		r := m.At(x+1, y) == palette[0]
		u := m.At(x, y-1) == palette[0]
		d := m.At(x, y+1) == palette[0]

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

func canvasIsFull(m image.RGBA, w int, h int, background color.RGBA) bool {
	cornerCount := 0

	if m.At(0, 0) != background {
		cornerCount++
	}

	if m.At(0, h) != background {
		cornerCount++
	}

	if m.At(w, 0) != background {
		cornerCount++
	}

	if m.At(w, h) != background {
		cornerCount++
	}

	if cornerCount == 4 {
		return true
	}

	return false
}
