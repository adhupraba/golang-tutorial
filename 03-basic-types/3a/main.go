package main

import "fmt"

func main() {
	a := 2
	b := 3.1

	fmt.Printf("a = %v -> %T\n", a, a)
	fmt.Printf("b = %v -> %8T\n\n", b, b)

	// %8T -> 8 is used to format aligment
	fmt.Printf("a = %v -> %8T\n", a, a)
	fmt.Printf("b = %v -> %8T\n\n", b, b)

	// %8T -> 8 is used to format aligment
	// [1] -> tells go to re-use the parameter 1 so we don't need to pass the same variable again
	// [0] is the string itself
	fmt.Printf("a = %v -> %8[1]T\n", a)
	fmt.Printf("b = %v -> %8[1]T\n\n", b)

	// float can't be assigned to int
	// we have to typecast and then assign
	a = int(b)
	fmt.Printf("a = %v -> %8[1]T\n", a)

}
