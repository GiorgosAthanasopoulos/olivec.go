package main

import (
	"./olivec_go"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

const WIDTH = 900
const HEIGHT = 600

var pixels []uint32

func main() {
	oc := olivec_go.OlivecCanvasNew(&pixels, WIDTH, HEIGHT, WIDTH)
	olivec_go.OlivecFill(oc, 0xFFFFFFFF)
	olivec_go.OlivecCircle(oc, WIDTH/2, HEIGHT/2, 180, 0xFF2D00BC)

	palette := make([]color.Color, len(*oc.pixels))
	for i, val := range *oc.pixels {
		palette[i] = color.RGBA{
			R: uint8(val >> 24),
			G: uint8(val >> 16),
			B: uint8(val >> 8),
			A: uint8(val),
		}
	}

	// Create a new image
	width := len(palette)
	height := 1
	img := image.NewPaletted(image.Rect(0, 0, width, height), palette)

	// Set pixel values
	for x := 0; x < width; x++ {

		img.Set(x, 0, palette[x])
	}

	// Create output file
	file, err := os.Create("output.png")
	if err != nil {
		panic(err)

	}
	defer file.Close()

	// Write image to file
	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}

	// Output success message
	fmt.Println("Image written successfully.")
}
