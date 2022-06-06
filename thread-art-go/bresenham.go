package main

import (
	"math"
)

type Point struct {
	X int
	Y int
}

// courtesy of fputs with ISC License
// https://github.com/fputs/bresenham/blob/3ab2d5f17f53/bresenham.go
// code has been slightly altered to only use standard library
func Bresenham(x0, y0, x1, y1 int) []Point {
	var line []Point
	var cx, cy, dx, dy, sx, sy, err int

	cx = x0
	cy = y0
	dx = int(math.Abs(float64(x1) - float64(x0)))
	dy = int(math.Abs(float64(y1) - float64(y0)))

	if cx < x1 {
		sx = 1
	} else {
		sx = -1
	}
	if cy < y1 {
		sy = 1
	} else {
		sy = -1
	}
	err = dx - dy

	for {
		line = append(line, Point{cx, cy})
		if cx == x1 && cy == y1 {
			return line
		}
		e2 := 2 * err
		if e2 > 0-dy {
			err = err - dy
			cx = cx + sx
		}
		if e2 < dx {
			err = err + dx
			cy = cy + sy
		}
	}
}
