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

type album1 struct {
	title string
}

type album2 struct {
	title string
}

func main() {
	var a1 = struct{ title string }{
		title: "Whitebeard",
	}

	var a2 = struct{ title string }{
		title: "Blackbeard",
	}

	a1 = a2 // ! will work since anonymous struct just checks for structural compatibility

	fmt.Println(a1, a2)

	var b1 = album1{
		title: "Whitebeard",
	}

	var b2 = album2{
		title: "Blackbeard",
	}

	// b1 = b2 // ! will not work since named types are different
	b1 = album1(b2) // ! will work because of explicit type conversion

	fmt.Println(b1, b2)

}
