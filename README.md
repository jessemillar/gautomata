> The cellular automaton consists of a line of cells, each colored either black or white. At every step there is then a definite rule that determines the color of a given cell from the color of that cell and its immediate left and right neighbors on the step before.
>
> -Stephen Wolfram (A New Kind of Science, 2004)

## Overview
`gautomata` is a [cellular automata](https://en.wikipedia.org/wiki/Cellular_automaton) renderer meant to generate wallpaper-size images of various cellular automata. Some of the available automata are well outside the simplistic bounds described by Stephen Worlfram's quote above, but that doesn't make them any less cellular and/or fun.

`gautomata` is written in the [Go programming language](https://golang.org/) to keep my Go skills sharp.

This project was heavily inspired by [trasevol_dog's](https://twitter.com/TRASEVOL_DOG) [writings on](https://trasevol.dog/2017/03/14/doodle-insights-8-cellular-automata-aka-black-magic/) and [experiments with](https://www.lexaloffle.com/bbs/?tid=28308) cellular automata.

## Sample
![Sample](https://github.com/jessemillar/gautomata/raw/master/automata.png)

## Installation
```
go get github.com/jessemillar/gautomata
```

## Usage
```
usage: gautomata [<flags>]

Flags:
      --help                   Show context-sensitive help (also try --help-long and --help-man).
  -w, --width=256              The width of the resulting image.
  -h, --height=256             The height of the resulting image.
  -c, --colors=5               The number of colors in the generated color palette.
  -l, --list                   List the supported automata.
  -a, --automata="random"      The automata to execute.
  -b, --batch                  Whether or not to automatically name the resulting images.
  -o, --output="automata.png"  The filename of the resulting image.
```

## Tips
`gautomata` is not capable of generating scaled images. If you wish to more easily see pixels in your resulting images, combining `gautomata` with something like [ImageMagick](https://www.imagemagick.org/script/index.php) could be useful. 

The following `bash`/[`zsh`](http://ohmyz.sh/) function uses `gautomata` to generate automata images and uses ImageMagick to scale them to MacBook Pro screen resolutions.
```
cells () {
        if [ -z $1 ]
        then
                echo "No iteration count supplied"
        else
                for i in {1..$1}
                do
                        if [ -z $2 ]
                        then
                                gautomata -w 576 -h 360 -b -a random
                        else
                                gautomata -w 576 -h 360 -b -a $2
                        fi
                done
                echo "Scaling images"
                mogrify -scale 5760x3600+0+0 *.png
        fi
}

```

## Notes
This repository uses [`gvt`](https://github.com/FiloSottile/gvt) for package management. Below is a list of dependencies and their versions that can be found in the `/vendor` directory.
```
gvt fetch -tag v2.2.5 gopkg.in/alecthomas/kingpin.v2
```
