package main

import (
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
	"gopkg.in/alecthomas/kingpin.v2"
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
	w := kingpin.Arg("width", "The width of the resulting image").Default("256").Int()
	h := kingpin.Arg("height", "The height of the resulting image").Default("256").Int()
	c := kingpin.Arg("palette", "The number of colors in the generated color palette").Default("6").Int()
	aut := kingpin.Arg("automata", "The automata to execute").Default("random").String()
	batch := kingpin.Flag("batch", "Whether or not to automatically name the resulting images").Short('b').Default("false").Bool()
	output := kingpin.Arg("output", "The filename of the resulting image").Default("automata.png").String()
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	// Generate a random color palette
	palette, err := tools.RandPalette(*c)
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
			log.Println("Started generating " + *aut)
			automata[k](*m, *w, *h, palette)
			break
		}
	} else { // Draw the selected automata
		log.Println("Started generating " + *aut)
		automata[*aut](*m, *w, *h, palette)
	}

	// Make the file
	if *batch { // Automatically name the output files if running as a batch
		*output = strconv.FormatInt(time.Now().UnixNano()/1000000, 10) + ".png"
	}

	f, err := os.OpenFile(*output, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	png.Encode(f, m)

	log.Println("Finished generating " + *aut)
}
