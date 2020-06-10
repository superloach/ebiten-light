package main

import "math"

func lineLow(x0, y0, x1, y1 float64) []Point {
	dx := x1 - x0
	dy := y1 - y0
	yi := 1.0
	if dy < 0.0 {
		yi = -1.0
		dy = -dy
	}
	D := 2*dy - dx
	y := y0

	pts := make([]Point, 0)
	for x := x0; x < x1; x++ {
		pts = append(pts, Point{int(x), int(y)})

		if D > 0 {
			y = y + yi
			D = D - 2*dx
		}
		D = D + 2*dy
	}

	return pts
}

func lineHigh(x0, y0, x1, y1 float64) []Point {
	dx := x1 - x0
	dy := y1 - y0
	xi := 1.0
	if dx < 0.0 {
		xi = -1.0
		dx = -dx
	}
	D := 2*dx - dy
	x := x0

	pts := make([]Point, 0)
	for y := y0; y < y1; y++ {
		pts = append(pts, Point{int(x), int(y)})

		if D > 0 {
			x = x + xi
			D = D - 2*dy
		}
		D = D + 2*dx
	}

	return pts
}

func line(a, b Point) []Point {
	x0 := float64(a[0])
	y0 := float64(a[1])
	x1 := float64(b[0])
	y1 := float64(b[1])

//	if y1 < y0 {
//		y0, y1 = y1, y0
//	}

	if math.Abs(y1-y0) < math.Abs(x1-x0) {
		if x0 > x1 {
			return lineLow(x1, y1, x0, y0)
		} else {
			return lineLow(x0, y0, x1, y1)
		}
	} else {
		if y0 > y1 {
			return lineHigh(x1, y1, x0, y0)
		} else {
			return lineHigh(x0, y0, x1, y1)
		}
	}
}
