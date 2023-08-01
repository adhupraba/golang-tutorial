package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

type ColoredPoint struct {
	// ! field name is not there, but type is there
	// ! the type becomes embeeded within the struct
	// ! here, fields of Point get promoted into ColoredPoint
	// ! (fields and methods of Point appear at the same level as the fields of ColoredPoint)
	Point
	Color color.RGBA
}

// ! a method may take a pointer or a value but not both

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	p, q := Point{1, 1}, ColoredPoint{Point{5, 4}, color.RGBA{255, 0, 0, 255}}

	// ! q.Distance is accessible because the Point type is embedded within ColoredPoint type
	d1 := q.Distance(p)
	d2 := p.Distance(q.Point)

	fmt.Println(d1)
	fmt.Println(d2)
}
