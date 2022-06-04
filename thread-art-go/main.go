package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"log"
	"math"
	"os"
)

var imagePath = "../selfie-exposure.jpg"
var numberOfPoints = 200
var rap = 1000
var minimumDifference = 20

// var minimumDifference = int(math.Floor(float64(numberOfPoints) / 10.0))

var brightnessFactor = 50

// func drawLine(pair string) {

// }

func getPair(a, b int) string {
	switch {
	case a < b:
		return fmt.Sprintf("%d-%d", a, b)
	case a > b:
		return fmt.Sprintf("%d-%d", b, a)
	default:
		fmt.Println("An error has occured - Please try again.")
		fmt.Printf("%d-%d\n", a, b)
		fmt.Printf("a=%d and b=%d\n", a, b)
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

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
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

	decodedImageConfig, _, err := image.DecodeConfig(f)

	if err != nil {
		fmt.Printf("Does file exist at %s with correct suffix\n", imagePath)
		fmt.Println(err.Error())
		os.Exit(1)
	}

	width := decodedImageConfig.Width
	height := decodedImageConfig.Height

	if width != 400 || height != 400 {
		fmt.Printf("Image at '%v' dimensions not supported.\n", imagePath)
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

	log.Printf("Image type: %T", decodedImage)

	// Convert image to gray
	// Makes it easier to get the pixel value
	// decodedImage is read only, but the new converted
	// image can be changed.
	bounds := decodedImage.Bounds()
	grayImage := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			grayImage.Set(x, y, decodedImage.At(x, y))
		}
	}

	// Get the sum of every pixel's red channel.
	// Test if it is the same as other languages.
	// Result: It is not the same as Python or Rust.
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

	var pointIndex = 0
	var pointsList = []int{pointIndex}
	var usedPairs map[string]struct{} = make(map[string]struct{})

	for i := 0; i < rap; i++ {
		var maxWeight = 0
		var maxLine = []Point{{0, 0}, {0, 0}}
		var maxPointIndex = 0

		//go through every line, calculate its weight and determine next point
		for nextPointIndex := 0; nextPointIndex < len(circle); nextPointIndex++ {
			if pointIndex == nextPointIndex {
				continue
			}

			var difference int = int(math.Abs(float64(nextPointIndex) - float64(pointIndex)))

			if difference < minimumDifference || difference > (len(circle)-minimumDifference) {
				continue
			}

			if _, exists := usedPairs[getPair(nextPointIndex, pointIndex)]; exists {
				continue
			}

			var line = lines[getPair(pointIndex, nextPointIndex)]
			var weight = len(line) * 255

			for _, pixelPosition := range line {
				pixel := grayImage.GrayAt(pixelPosition.X, pixelPosition.Y)
				// Value of pixel goes from 0 to 255
				value := pixel.Y
				weight -= int(value)
			}

			weight = weight / len(line)

			if weight > maxWeight {
				maxLine = line
				maxPointIndex = nextPointIndex
			}
		}

		usedPairs[getPair(pointIndex, maxPointIndex)] = struct{}{}
		pointsList = append(pointsList, maxPointIndex)
		pointIndex = maxPointIndex

		// Brighthen brightness of chosen line
		for _, pixelPosition := range maxLine {
			var pixel = int(grayImage.GrayAt(pixelPosition.X, pixelPosition.Y).Y)
			value := uint8(min(255, pixel+brightnessFactor))
			grayImage.SetGray(pixelPosition.X, pixelPosition.Y, color.Gray{value})
		}
	}

	fmt.Printf("Totals to %v different lines = n*(n-1)/2\n", len(lines))
	fmt.Printf("Total pointsList = %v\n", len(pointsList))
	fmt.Printf("\n%v\n", pointsList)
	fmt.Printf("Total usedPairs = %v\n", len(usedPairs))

	drawImage := image.NewGray(bounds)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			drawImage.SetGray(x, y, color.Gray{255})
		}
	}

	// draw line
	for i := 1; i < len(pointsList); i++ {
		firstPointIndex := pointsList[i-1]
		secondPointIndex := pointsList[i]
		line := lines[getPair(firstPointIndex, secondPointIndex)]
		for _, point := range line {
			currentValue := drawImage.GrayAt(point.X, point.Y).Y
			newValue := max(int(currentValue)-20, 0)
			drawImage.SetGray(point.X, point.Y, color.Gray{uint8(newValue)})
		}
	}

	// save image as png
	outFile, err := os.Create("outimage.png")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = png.Encode(outFile, drawImage)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	outFile.Close()
}
