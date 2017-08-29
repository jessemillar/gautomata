package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

var (
	white color.Color = color.RGBA{255, 255, 255, 255}
	black color.Color = color.RGBA{0, 0, 0, 255}
	blue  color.Color = color.RGBA{0, 0, 255, 255}
)

func main() {
	// Get params
	w := *flag.Int("w", 256, "the width of the resulting image")
	h := *flag.Int("h", 256, "the height of the resulting image")
	output := *flag.String("o", "automata.png", "the filename of the resulting image")
	flag.Parse()

	m := image.NewRGBA(image.Rect(0, 0, w, h))
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src) // Fill with a uniform color

	// Drawing happens here
	m.Set(10, 10, white)
	fmt.Println(m.At(0, 0))
	fmt.Println(m.At(10, 10))

	// Make the file
	f, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	png.Encode(f, m)
}

/*
for i=0,127 do pset(i,0,flr(rnd(2))) end
--random initial states ^^^
palt(0,false)

function _draw()
 memcpy(0x0,0x6000,0x2000)
 spr(0,0,1,16,16)
 line(0,0,127,0,0)

 for x=0,127 do
  local k=0

  k+=pget(x-1,1)
  k+=pget(x+1,1)

  if k==1 then
   pset(x,0,(k)%2)
  end
 end
end
*/
