package main

import (
	"fmt"
	"time"
)

type Employee struct {
	Name   string
	Number int
	Boss   *Employee
	Hired  time.Time
}

func main() {
	e := Employee{
		Name:   "Matt",
		Number: 1,
		Hired:  time.Now(),
	}

	b := Employee{
		Name:   "Lamine",
		Number: 2,
		Hired:  time.Now(),
	}

	e.Boss = &b

	fmt.Printf("%T %+[1]v\n", e)
	fmt.Printf("%T %+[1]v\n", b)
}
