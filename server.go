package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"os"
	"time"
)

var (
	dark  color.Color = color.RGBA{68, 71, 90, 255}
	light color.Color = color.RGBA{98, 114, 164, 255}
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	// Get params
	w := flag.Int("w", 256, "the width of the resulting image")
	h := flag.Int("h", 256, "the height of the resulting image")
	output := flag.String("o", "automata.png", "the filename of the resulting image")
	flag.Parse()

	m := image.NewRGBA(image.Rect(0, 0, *w, *h))
	draw.Draw(m, m.Bounds(), &image.Uniform{dark}, image.ZP, draw.Src) // Fill with a uniform color

	// Random initial state
	for i := 0; i < *w; i++ {
		if rand.Intn(2) < 1 {
			m.Set(i, 0, light)
		}
	}

	// Loop through the y axis
	for y := 1; y < *h; y++ {
		for x := 0; x < *w; x++ {
			draw := 0

			if m.At(x-1, y-1) == dark {
				draw++
			}

			if m.At(x+1, y-1) == dark {
				draw++
			}

			if draw == 1 {
				m.Set(x, y, light)
			}
		}
	}

	// Make the file
	f, err := os.OpenFile(*output, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	png.Encode(f, m)
}
