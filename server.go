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
	"strconv"
	"strings"
	"time"

	"github.com/jessemillar/gautomata/cells"
	"github.com/jessemillar/gautomata/tools"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	// Get params
	w := flag.Int("w", 256, "the width of the resulting image")
	h := flag.Int("h", 256, "the height of the resulting image")
	col := flag.String("c", "#ffffff", "the color to use as an average for randomly-generated colors")
	output := flag.String("o", "automata.png", "the filename of the resulting image")
	flag.Parse()

	// Parse the mix color
	colHex := strings.TrimLeft(*col, "#") // Strip the #
	fmt.Println(colHex)
	colRed, _ := strconv.ParseUint(colHex[0:2], 16, 32)
	colGreen, _ := strconv.ParseUint(colHex[2:4], 16, 32)
	colBlue, _ := strconv.ParseUint(colHex[4:6], 16, 32)
	mix := color.RGBA{uint8(colRed), uint8(colGreen), uint8(colBlue), 255}

	// Generate random colors
	light := tools.RandColor(mix)
	dark := tools.RandColor(mix)

	// Make an image
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
