package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	svg "github.com/ajstarks/svgo"
)

func rn(n int) int {
	return rand.Intn(n)
}

func main() {
	f, err := os.OpenFile("output.svg", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer f.Close()

	//data:=[]int{100,33,73,64}
	canvas := svg.New(f)
	width, height, nstars := 500, 500, 250
	style := "font-size:48pt;fill:white;text-anchor:middle"

	rand.Seed(time.Now().Unix())
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height)
	for range nstars {
		canvas.Circle(rn(width), rn(height), rn(3), "fill:white")
	}
	canvas.Circle(width/2, height/2, 100, "fill:red")
	canvas.Text(width/2, height/2, "Hello World", style)
	canvas.End()
}
