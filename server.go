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

	"github.com/jessemillar/gautomata/cells"
	"github.com/jessemillar/gautomata/tools"
)

var mix color.Color = color.RGBA{255, 255, 255, 255}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	// Get params
	w := flag.Int("w", 256, "the width of the resulting image")
	h := flag.Int("h", 256, "the height of the resulting image")
	output := flag.String("o", "automata.png", "the filename of the resulting image")
	flag.Parse()

	light := tools.RandColor(mix)
	dark := tools.RandColor(mix)

	m := image.NewRGBA(image.Rect(0, 0, *w, *h))
	draw.Draw(m, m.Bounds(), &image.Uniform{dark}, image.ZP, draw.Src) // Fill with a uniform color

	// Draw the automata
	// cells.Funnels(*m, *w, *h, light, dark)
	// cells.Rule30(*m, *w, *h, light, dark)
	cells.Rule110(*m, *w, *h, light, dark)

	// Make the file
	f, err := os.OpenFile(*output, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	png.Encode(f, m)
}
