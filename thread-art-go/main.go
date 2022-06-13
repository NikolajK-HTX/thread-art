package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"log"
	"math"
	"os"
	"time"
)

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
	log.SetFlags(0)
	log.SetPrefix("Log: ")

	start := time.Now()

	var imagePath string
	var outputName string
	var numberOfPins int
	var rap int
	var minimumDifference int
	// var minimumDifference = int(math.Floor(float64(numberOfPoints) / 10.0))
	var brightnessFactor int
	var lineAlpha int

	flag.StringVar(&imagePath, "i", "../selfie-exposure.jpg",
		"The path to the image.")
	flag.StringVar(&outputName, "o", "output-image",
		"Output filename. Do not include suffix.")
	flag.IntVar(&numberOfPins, "n", 200,
		"A higher amount of pins makes the result more precise.")
	flag.IntVar(&rap, "t", 3000, "Amount of thread raps")
	flag.IntVar(&minimumDifference, "d", 20,
		"The next pin can't be next to current pin.")
	flag.IntVar(&brightnessFactor, "b", 50,
		"Amount to brighten pixel")
	flag.IntVar(&lineAlpha, "l", 20, "The alpha of the drawn lines.")
	flag.Parse()

	end := time.Now()
	elapsed := end.Sub(start)
	log.Printf("Time spent on parsing flags %.2f\n", elapsed.Seconds())

	start = time.Now()
	f, _ := os.Open(imagePath)

	// using image.DecodeConfig(f) consumes all of f
	// so you need to seek to (0, 0) if open the file again.
	// https://stackoverflow.com/questions/62846156/image-decode-unknown-format
	// the benefit is that the image is not read so it is
	// possible to check the dimensions before loading

	decodedImageConfig, _, err := image.DecodeConfig(f)

	if err != nil {
		log.Printf("Does file exist at %s with correct suffix\n", imagePath)
		log.Fatalln(err.Error())
	}

	width := decodedImageConfig.Width
	height := decodedImageConfig.Height

	if width != 400 || height != 400 {
		log.Printf("Image at '%v' dimensions not supported.\n", imagePath)
		log.Fatalln("Only square by 400 pixels is accepted.")
	}

	f.Seek(0, 0)

	decodedImage, myString, err := image.Decode(f)

	if err != nil {
		log.Println(myString)
		log.Fatalln(err.Error())
	}

	// log.Printf("Image type: %T", decodedImage)

	// Convert image to gray. Makes it easier to get the pixel value
	// decodedImage is read only, but the new converted image can be changed.
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
	for i := 0; i < numberOfPins; i++ {
		var x, y int
		angle := float64(math.Pi) * 2.0 / float64(numberOfPins) * float64(i)
		x = constrain(int(math.Cos(angle)*float64(width)/2.0+float64(width)/2.0), 0, width)
		y = constrain(int(math.Sin(angle)*float64(height)/2.0+float64(height)/2.0), 0, width)
		circle = append(circle, Point{x, y})
	}

	// create a dictionary of all possible lines in circle.
	// it is more performant to create this once and have the
	// results at hand than creating them when needed to check
	// the weight of possible lines and decide which to draw a line on.

	var lines map[string][]Point = make(map[string][]Point)

	for a := 0; a < numberOfPins; a++ {
		for b := a + 1; b < numberOfPins; b++ {
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

			difference := int(math.Abs(float64(nextPointIndex) - float64(pointIndex)))

			if difference < minimumDifference || difference > (len(circle)-minimumDifference) {
				continue
			}

			if _, exists := usedPairs[getPair(nextPointIndex, pointIndex)]; exists {
				continue
			}

			line := lines[getPair(pointIndex, nextPointIndex)]
			weight := len(line) * 255

			for _, pixelPosition := range line {
				pixelColor := grayImage.GrayAt(pixelPosition.X, pixelPosition.Y).Y
				weight -= int(pixelColor)
			}

			weight = weight / len(line)

			if weight > maxWeight {
				maxWeight = weight
				maxLine = line
				maxPointIndex = nextPointIndex
			}
		}

		usedPairs[getPair(pointIndex, maxPointIndex)] = struct{}{}
		pointsList = append(pointsList, maxPointIndex)
		pointIndex = maxPointIndex

		// Brighthen pixel of chosen line
		for _, pixelPosition := range maxLine {
			var pixel = int(grayImage.GrayAt(pixelPosition.X, pixelPosition.Y).Y)
			value := uint8(min(255, pixel+brightnessFactor))
			grayImage.SetGray(pixelPosition.X, pixelPosition.Y, color.Gray{value})
		}
	}

	// log.Printf("Totals to %v different lines = n*(n-1)/2\n", len(lines))
	// log.Printf("Total pointsList = %v\n", len(pointsList))
	// log.Printf("Total usedPairs = %v\n", len(usedPairs))

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
			newValue := max(int(currentValue)-lineAlpha, 0)
			drawImage.SetGray(point.X, point.Y, color.Gray{uint8(newValue)})
		}
	}

	// save image as png
	outFile, err := os.Create(fmt.Sprintf("%s.png", outputName))
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = png.Encode(outFile, drawImage)
	if err != nil {
		log.Fatalln(err.Error())
	}

	outFile.Close()

	// Write the list of pins the thread is going around
	outPointFile, err := os.Create(
		fmt.Sprintf("%s-thread-pin-list.txt", outputName))
	if err != nil {
		log.Fatalln(err.Error())
	}

	outString := ""

	for _, pointIndex := range pointsList {
		outString += fmt.Sprintf("%d,\n", pointIndex)
	}

	outPointFile.WriteString(outString)
	outPointFile.Close()

	end = time.Now()
	elapsed = end.Sub(start)
	log.Printf("Time spent on algorithm %.2f\n", elapsed.Seconds())

	// unsafe.Sizeof returns the size in bytes of a hypothetical variable
	// "For instance, if x is a slice, Sizeof returns the size of the
	// slice descriptor, not the size of the memory referenced by the slice."
	// - Go documentation
	// It is not what is wanted here.
	// log.Printf("decodedImage memory: %v", unsafe.Sizeof(decodedImage))
	// log.Printf("grayImage memory: %v", unsafe.Sizeof(grayImage))
	// log.Printf("drawImage memory: %v", unsafe.Sizeof(drawImage))
	// log.Printf("Circle memory: %v", unsafe.Sizeof(circle))
	// log.Printf("Lines memory: %v", unsafe.Sizeof(lines))
	// The following line gives a way to tell how many bytes of memory
	// a variable is using.
	// v, _ := json.Marshal(lines)
	// log.Println(len(v))
}
