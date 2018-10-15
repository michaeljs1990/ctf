package main

import (
	"bytes"
	"flag"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"os"
	"strconv"
)

var logger *log.Logger

func main() {
	var buf bytes.Buffer
	logger = log.New(&buf, "log: ", log.Lshortfile)

	imgFlag := flag.String("img", "", "Image that you would like to load")
	flag.Parse()

	// switch arg := os.Args[1]; arg {
	// case "alpha":
	// default:
	// }

	dat := loadImage(*imgFlag)
	aImg, _ := png.Decode(dat)

	// Make a empty image that can hold the input image
	nImg := image.NewRGBA(aImg.Bounds())
	for i := 0; i <= 1; i++ {
		for y := aImg.Bounds().Min.Y; y < aImg.Bounds().Max.Y; y++ {
			for x := aImg.Bounds().Min.X; x < aImg.Bounds().Max.X; x++ {
				c := aImg.At(x, y)
				_, _, _, a := c.RGBA()
				//ii := uint(i)
				// a = ((a >> ii) & 1)
				// if a > 0 {
				// 	a = 255
				// }
				if ((a >> 0)&1) > 0 {
					a = 0xff
				}
				col := color.RGBA{
					R: uint8(255),
					G: uint8(255),
					B: uint8(255),
					A: uint8(a),
				}
				nImg.Set(x, y, col)
			}
		}

		f, err := os.Create("out/test" + strconv.Itoa(i) + ".png")
		if err != nil {
			log.Fatal(err)
		}

		if err := png.Encode(f, nImg); err != nil {
			f.Close()
			log.Fatal(err)
		}

		if err := f.Close(); err != nil {
			log.Fatal(err)
		}

	}
	// fmt.Printf("%v", aImg)
}

func loadImage(name string) *os.File {
	dat, err := os.Open(name)
	if err != nil {
		log.Fatal("Unable to open file got error: ", err)
	}

	return dat
}

func writeImage(name string, r io.Reader) {

}
