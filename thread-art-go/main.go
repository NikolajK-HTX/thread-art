package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"log"
	"math"
	"os"
)

var imagePath = "../selfie.jpg"
var numberOfPoints = 200
var rap = 1000
var minimumDifference int = int(math.Floor(float64(numberOfPoints) / 10.0))

func getPair(a, b int) string {
	switch {
	case a < b:
		return fmt.Sprintf("%d-%d", a, b)
	case a > b:
		return fmt.Sprintf("%d-%d", b, a)
	default:
		fmt.Println("An error has occured - Please try again.")
		os.Exit(1)
		return ""
	}
}

func constrain(x, a, b int) int {
	switch {
	case x < a:
		return a
	case x > b:
		return b
	default:
		return x
	}
}

func main() {
	f, _ := os.Open(imagePath)

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
		fmt.Printf("Image at '%v' dimensions not supported.\n", imagePath)
		fmt.Println("Only square by 400 pixels is accepted.")

		os.Exit(1)
	}

	f.Seek(0, 0)

	decodedImage, myString, err := image.Decode(f)

	log.Printf("Image type: %T", decodedImage)

	if err != nil {
		fmt.Println(myString)
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// var sum uint64 = 0
	// for x := 1; x < width; x++ {
	// 	for y := 1; y < height; y++ {
	// 		pixel := decodedImage.At(x, y)
	// 		R, _, _, _ := pixel.RGBA()

	// 		sum += uint64(R / 257)
	// 	}
	// }

	// create points outlining circle
	var circle []Point
	for i := 0; i < numberOfPoints; i++ {
		var x, y int
		angle := float64(math.Pi) * 2.0 / float64(numberOfPoints) * float64(i)
		x = constrain(int(math.Cos(angle)*float64(width)/2.0+float64(width)/2.0), 0, width)
		y = constrain(int(math.Sin(angle)*float64(height)/2.0+float64(height)/2.0), 0, width)
		circle = append(circle, Point{x, y})
	}

	// create a dictionary of all possible lines in circle.
	// it is more performant to create this once and have the
	// results at hand than creating them when needed to check
	// the weight of possible lines and decide which to draw a line on.

	var lines map[string][]Point = make(map[string][]Point)

	for a := 0; a < numberOfPoints; a++ {
		for b := a + 1; b < numberOfPoints; b++ {
			var pair = fmt.Sprintf("%d-%d", a, b)
			var x0, y0, x1, y1 int

			x0 = circle[a].X
			y0 = circle[a].Y
			x1 = circle[b].X
			y1 = circle[b].Y

			lines[pair] = Bresenham(x0, y0, x1, y1)
		}
	}

	var startPointIndex = 0
	var _ = circle[startPointIndex]

	for i := 0; i < rap; i++ {
		var maxWeight = 0
		var maxLine = []Point{{0, 0}, {0, 0}}

		//go through every line and determine weight
		for index, _ := range circle {
			var difference int = int(math.Abs(float64(index) - float64(startPointIndex)))
			if difference < minimumDifference || difference > (len(circle)-minimumDifference) {
				continue
			}
			var line = lines[getPair(startPointIndex, index)]
			var weight int = len(line) * 255
			for _, pixelPosition := range line {
				pixel := decodedImage.At(pixelPosition.X, pixelPosition.Y)
				// Value of pixel goes from 0 to 255
				value := color.GrayModel.Convert(pixel).(color.Gray).Y
				weight -= int(value)
			}

			if weight > maxWeight {
				maxLine = line
			}
		}
	}

	fmt.Printf("Totals to %v different lines = n*(n-1)/2\n", len(lines))
}
