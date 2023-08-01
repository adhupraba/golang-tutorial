package main

import "fmt"

func fib() func() int {
	a, b := 0, 1

	return func() int {
		a, b = b, a+b
		return b
	}
}

/*
	a		b
	0		1
	1		1
	1		2
	2		3
	3		5
	5		8
	8		13
	.		.
	.		.
	.		.
*/

func main() {
	f := fib()

	for x := f(); x < 100; x = f() {
		fmt.Println(x)
	}
}
