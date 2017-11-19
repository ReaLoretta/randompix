package main

import (
	"fmt"
	"image"
	"image/color"
	"os"

	"golang.org/x/image/bmp"
)

func main() {
	// Create image
	size := 128
	img := image.NewRGBA(image.Rect(0, 0, size, size))

	// Populate image
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			img.Set(i, j, color.RGBA{
				getRandomByte(),
				getRandomByte(),
				getRandomByte(),
				getRandomByte(),
			})
		}
	}

	// Save to random.bmp
	f, _ := os.OpenFile("random.bmp", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	bmp.Encode(f, img)

	fmt.Printf("Bits left: %d\n", checkQuota())
}
