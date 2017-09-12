package cells

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/jessemillar/gautomata/tools"
)

func Flowers(m image.RGBA, w int, h int, background color.Color) {
	foreground := tools.RandColor()

	// tools.RandState(m, w, foreground)
	m.Set(w/2, h/2, foreground)

	for !canvasIsFull(m, w, h, background) {
		x := rand.Intn(w)
		y := rand.Intn(h)

		if m.At(x, y) == background {
			continue
		}

		if rand.Intn(2) < 1 {
			foreground = tools.RandColor()
		}

		l := m.At(x-1, y) == background
		r := m.At(x+1, y) == background
		u := m.At(x, y-1) == background
		d := m.At(x, y+1) == background

		if l {
			m.Set(x-1, y, foreground)
		}

		if r {
			m.Set(x+1, y, foreground)
		}

		if u {
			m.Set(x, y-1, foreground)
		}

		if d {
			m.Set(x, y+1, foreground)
		}
	}
}

func canvasIsFull(m image.RGBA, w int, h int, background color.Color) bool {
	/*
		// Loop through the canvas to check if every pixel is colored
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				if m.At(x, y) == background {
					return false
				}
			}
		}
	*/

	if m.At(0, 0) == background {
		return false
	}

	if m.At(0, h) == background {
		return false
	}

	if m.At(w, 0) == background {
		return false
	}

	if m.At(w, h) == background {
		return false
	}

	return true
}
