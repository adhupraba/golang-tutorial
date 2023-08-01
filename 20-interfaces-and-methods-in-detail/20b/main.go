package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	p := Point{1, 1}
	q := Point{5, 4}

	fmt.Println(p.Distance(q))

	distanceFromP := p.Distance
	fmt.Printf("%T\n", distanceFromP)
	fmt.Printf("%T\n", Point.Distance)

	p = Point{2, 2}
	fmt.Println(p.Distance(q))

	fmt.Println(distanceFromP(q))
}
