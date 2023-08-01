package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Line struct {
	Begin, End Point
}

type Path []Point

type Distancer interface {
	Distance() float64
}

// ! a method may take a pointer or a value but not both

func (l Line) Distance() float64 {
	return math.Hypot(l.End.X-l.Begin.X, l.End.Y-l.Begin.Y)
}

func (l *Line) ScaleBy(f float64) {
	l.End.X += (f - 1) * (l.End.X - l.Begin.X)
	l.End.Y += (f - 1) * (l.End.Y - l.Begin.Y)
}

func (p Path) Distance() (sum float64) {
	for i := 1; i < len(p); i++ {
		sum += Line{Begin: p[i-1], End: p[i]}.Distance()
	}

	return sum
}

func PrintDistance(d Distancer) {
	fmt.Println(d.Distance())
}

func main() {
	side := Line{Point{1, 2}, Point{4, 6}}
	// ! typing each value in the array with Point is redundant
	// perimeter := Path{Point{1, 1}, Point{5, 1}, Point{5, 4}, Point{1, 1}}
	perimeter := Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}

	side.ScaleBy(2)

	// fmt.Println(side.Distance())
	// fmt.Println(perimeter.Distance())

	PrintDistance(side)
	PrintDistance(perimeter)
}
