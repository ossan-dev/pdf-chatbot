package main

import (
	"image/png"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	imgFile, err := os.Open("imgs/booking.png")
	if err != nil {
		panic(err)
	}
	defer imgFile.Close()
	image, err := png.Decode(imgFile)
	if err != nil {
		panic(err)
	}

	resizedImage := resize.Resize(448, 448, image, resize.Lanczos3)

	imgResizedFile, err := os.Create("imgs/booking_resized.png")
	if err != nil {
		panic(err)
	}
	if err := png.Encode(imgResizedFile, resizedImage); err != nil {
		panic(err)
	}

}
