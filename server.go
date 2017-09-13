package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/jessemillar/gautomata/cells"
	"github.com/jessemillar/gautomata/tools"
)

var automata = map[string]func(image.RGBA, int, int, []color.RGBA){
	"funnels": cells.Funnels,
	"flowers": cells.Flowers,
	"rule30":  cells.Rule30,
	"rule110": cells.Rule110,
	"rule126": cells.Rule126,
	"rule150": cells.Rule150,
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano()) // Set the random seed

	// Get params
	w := flag.Int("w", 256, "the width of the resulting image")
	w = flag.Int("width", 256, "the width of the resulting image")
	h := flag.Int("h", 256, "the height of the resulting image")
	p := flag.Int("p", 6, "the number of colors in the generated color palette")
	aut := flag.String("a", "random", "the automata to execute")
	batch := flag.Bool("b", false, "whether or not to automatically name the resulting images")
	output := flag.String("o", "automata.png", "the filename of the resulting image")
	flag.Parse()

	// Generate a random color palette
	palette, err := tools.RandPalette(*p)
	if err != nil {
		log.Fatal(err)
	}

	// Make an image
	m := image.NewRGBA(image.Rect(0, 0, *w, *h))
	draw.Draw(m, m.Bounds(), &image.Uniform{palette[0]}, image.ZP, draw.Src) // Fill with a uniform color

	// Draw the automata
	if *aut == "random" { // Select a random automata
		for k, _ := range automata {
			*aut = k
			fmt.Println("Started generating " + *aut)
			automata[k](*m, *w, *h, palette)
			break
		}
	} else { // Draw the selected automata
		fmt.Println("Started generating " + *aut)
		automata[*aut](*m, *w, *h, palette)
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
