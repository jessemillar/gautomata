package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

type Circle struct {
	X, Y, R float64
}

func (c *Circle) Brightness(x, y float64) uint8 {
	var dx, dy float64 = c.X - x, c.Y - y
	d := math.Sqrt(dx*dx+dy*dy) / c.R
	if d > 1 {
		return 0
	} else {
		return 255
	}
}

func main() {
	w := *flag.Int("w", 256, "the width of the resulting image")
	h := *flag.Int("h", 256, "the height of the resulting image")
	flag.Parse()

	var hw, hh float64 = float64(w / 2), float64(h / 2)
	r := 40.0
	θ := 2 * math.Pi / 3
	cr := &Circle{hw - r*math.Sin(0), hh - r*math.Cos(0), 60}
	cg := &Circle{hw - r*math.Sin(θ), hh - r*math.Cos(θ), 60}
	cb := &Circle{hw - r*math.Sin(-θ), hh - r*math.Cos(-θ), 60}

	m := image.NewRGBA(image.Rect(0, 0, w, h))
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			c := color.RGBA{
				cr.Brightness(float64(x), float64(y)),
				cg.Brightness(float64(x), float64(y)),
				cb.Brightness(float64(x), float64(y)),
				255,
			}
			m.Set(x, y, c)
		}
	}

	f, err := os.OpenFile("rgb.png", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	png.Encode(f, m)
}

/*
--palette and flower generation
--pico-8 doodle by trasevol_dog

--the used palette is generated
--from preset palettes in 'plts'



cls()

plts={
{1,2,4,9,10,15},
{15,10,9,4,2,1},
{11,3,1},
{7,6,13,1},
{7,6,13,12,11,10,9,4,2},
{15,10,9,8,14},
{14,8,2,13,1},
{7,14,8,2},
{7,15,9}
}

::palette::

k=flr(rnd(6))+3
p=plts[flr(rnd(#plts))+1]
c=p[flr(rnd(#p))+1]
plt={c}
for i=1,k do
 --pset(#plt,0,c)
 lc=c
 t=0
 while c==lc do
  t+=0.01
  p=plts[flr(rnd(#plts))+1]
  for j=1,#p do
   if p[j]==c then
    a=sgn(flr(rnd(2))-0.5)
    if(j==#p) a=-abs(a)
    if(j==1) a=abs(a)

    nc=p[j+a]
    good=true
    for l=1,#plt do
     good=(good and nc~=plt[l])
    end
    if(good)c=nc
   end
  end

  if(t>100) goto palette
 end
 add(plt,c)
end

--plt=plts[#plts]
cls()
pset(48+rnd(32),48+rnd(32),plt[1])

::s::

x=rnd(128)
y=rnd(128)

c=pget(x,y)

if(c==0)goto s

k=1
while c~=plt[k] do k+=1 end

if rnd(2)<1 then k=(k)%#plt+1 end

c=plt[k]

if(pget(x-1,y)==0)pset(x-1,y,c)
if(pget(x+1,y)==0)pset(x+1,y,c)
if(pget(x,y-1)==0)pset(x,y-1,c)
if(pget(x,y+1)==0)pset(x,y+1,c)

if(btn(4)) goto palette
goto s
*/
