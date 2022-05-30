package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"os"
)

func main() {
	fmt.Println("Hello!")
	f, _ := os.Open("../selfie.jpg")

	// using image.DecodeConfig(f) consumes all of f
	// so you need to seek to (0, 0) if open the file
	// again.
	// https://stackoverflow.com/questions/62846156/image-decode-unknown-format
	// the benefit is that the image is not read so it is
	// possible to check the dimensions before loading

	decodedImageConfig, _, _ := image.DecodeConfig(f)

	width := decodedImageConfig.Width
	height := decodedImageConfig.Height

	if width != 400 || height != 400 {
		fmt.Println("Image dimensions not supported.")
		fmt.Println("Only square by 400 pixels is accepted.")

		os.Exit(1)
	}

	f.Seek(0, 0)

	decodedImage, myString, err := image.Decode(f)

	if err != nil {
		fmt.Println(myString)
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var sum uint64 = 0
	for x := 1; x < width; x++ {
		for y := 1; y < height; y++ {
			pixel := decodedImage.At(x, y)
			R, _, _, _ := pixel.RGBA()

			sum += uint64(R / 257)
		}
	}
	fmt.Println(sum)

	// test bresenham
	fmt.Printf("%v", Bresenham(1, 1, 10, 10))
}
