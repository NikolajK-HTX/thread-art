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
	
	// newImageConfig, _, _ := image.DecodeConfig(f)

	// width := newImageConfig.Width
	// height := newImageConfig.Height

	// if width > 400 || height > 400 {
	// 	fmt.Println("Image dimensions not supported.")
	// 	fmt.Println("The limit is 400 pixels.")

	// 	os.Exit(1)
	// }

	newImage, myString, err := image.Decode(f)

	width := newImage.Bounds().Size().X
	height := newImage.Bounds().Size().Y
	
	if err != nil {
		fmt.Println(myString)
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var sum uint64 = 0
	for x := 1; x < width; x++ {
		for y := 1; y < height; y++ {
			pixel := newImage.At(x, y)
			R, _, _, _ := pixel.RGBA()

			sum += uint64(R/257)
		}
	}
	fmt.Println(sum)
}
