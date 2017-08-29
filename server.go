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
	"time"

	"github.com/jessemillar/gautomata/cells"
	"github.com/jessemillar/gautomata/tools"
)

var automata = map[string]func(image.RGBA, int, int, color.Color){
	"funnels": cells.Funnels,
	"rule30":  cells.Rule30,
	"rule110": cells.Rule110,
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano()) // Set the random seed

	// Get params
	w := flag.Int("w", 256, "the width of the resulting image")
	h := flag.Int("h", 256, "the height of the resulting image")
	aut := flag.String("a", "random", "the automata to execute")
	batch := flag.Bool("b", false, "whether or not to automatically name the resulting images")
	output := flag.String("o", "automata.png", "the filename of the resulting image")
	flag.Parse()

	// Generate the background color
	background := tools.RandColor()

	// Make an image
	m := image.NewRGBA(image.Rect(0, 0, *w, *h))
	draw.Draw(m, m.Bounds(), &image.Uniform{background}, image.ZP, draw.Src) // Fill with a uniform color

	// Draw the automata
	if *aut == "random" { // Select a random automata
		for k, _ := range automata {
			*aut = k
			automata[k](*m, *w, *h, background)
			break
		}
	} else { // Draw the selected automata
		automata[*aut](*m, *w, *h, background)
	}

	// Make the file
	if *batch { // Automatically name the output files if running as a batch
		*output = strconv.FormatInt(time.Now().UnixNano()/1000000, 10) + ".png"
	}

	f, err := os.OpenFile(*output, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	png.Encode(f, m)

	fmt.Println("Finished generating " + *aut)
}
