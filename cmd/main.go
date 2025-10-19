package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	data := []int{100, 33, 73, 64}
	w, h := len(data)*60+10, 100

	r := image.Rect(0, 0, w, h)
	img := image.NewRGBA(r)
	bg := image.NewUniform(color.RGBA{240, 240, 240, 255})
	draw.Draw(img, img.Bounds(), bg, image.Point{0, 0}, draw.Src)

	mask := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := range h {
		for x := range w {
			alpha := 0
			switch {
			case y >= 0 && y < 30:
				alpha = 255
			case y >= 30 && y < 60:
				alpha = 100
			}
			mask.Set(x, y, color.RGBA{
				R: uint8((x + y) & 255),
				G: uint8((x + y) << 1 & 255),
				B: uint8((x + y) << 1 & 255),
				A: uint8(alpha),
			})
		}
	}

	// for y := range h {
	// 	for x := range w {
	// 		img.Set(x, y, color.RGBA{255, 255, 255, 255})
	// 	}
	// }

	for i, dp := range data {
		x0, y0 := (i*60 + 10), 100-dp
		x1, y1 := (i+1)*60-1, 100
		bar := image.Rect(x0, y0, x1, y1)

		grey := image.NewUniform(color.RGBA{180, 180, 180, 255})
		draw.Draw(img, bar, grey, image.Point{0, 0}, draw.Src)

		red := image.NewUniform(color.RGBA{250, 180, 180, 255})
		draw.DrawMask(img, bar, red, image.Point{0, 0}, mask, image.Point{x0, y0}, draw.Over)

		// for x := i*60 + 10; x < i*60+60; x++ {
		// 	for y := 100; y >= (100 - dp); y-- {
		// 		img.Set(x, y, color.RGBA{180, 180, 250, 255})
		// 	}
		// }
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
