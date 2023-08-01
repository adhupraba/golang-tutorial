package main

import "fmt"

func main() {
	items := [][2]byte{{1, 2}, {3, 4}, {5, 6}}
	a := [][]byte{}
	b := [][]byte{}

	for _, item := range items {
		fmt.Println("item =>", item[:])
		// eventhough in each iteration "item" has its current iteration value
		// since it is a slice, when appending it into "a", the value is not copied into a
		// but the reference is pushed into "a" and the variable "item" has the same reference throughout the loop
		// when the loop finished, the value [5, 6] which is present in the last iteration will be printed when "a" is logged into console
		// since that is what item reference in the last iteration
		a = append(a, item[:])

		fmt.Println("a =>", a)

		// to tackle this issue, we create a new slice and copy the "item" data into the new slice
		// and append that new slice into "b"
		temp := make([]byte, len(item)) // this contains the current iteration value unlike the above code
		copy(temp, item[:])             // make unique
		b = append(b, temp)
	}

	fmt.Println(items)
	fmt.Println(a)
	fmt.Println(b)
}
