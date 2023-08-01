package main

import "fmt"

type Point struct {
	X, Y float64
}

// ! a method may take a pointer or a value but not both at the same time

func (p Point) Offset(x, y float64) Point {
	return Point{X: p.X + x, Y: p.Y + y}
}

func (p *Point) Move(x, y float64) {
	p.X += x
	p.Y += y
}

func main() {
	var p = Point{X: 2, Y: 5}

	var q = p.Offset(3, 7)
	fmt.Printf("p after offset => %v\n", p)
	p.Move(10, 5)

	fmt.Printf("p after move => %v\n", p)
	fmt.Printf("q => %v\n", q)
}
