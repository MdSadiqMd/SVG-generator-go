package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	w, h := 500, 150
	r := image.Rect(0, 0, w, h)
	img := image.NewRGBA(r)

	for y := range h {
		for x := range w {
			img.Set(x, y, color.RGBA{
				R: uint8((x + y) & 255),
				G: uint8((x + y) << 1 & 255),
				B: uint8((x + y) << 1 & 255),
				A: 255,
			})
		}
	}

	f, err := os.Create("image.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}
