package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3} // array
	b := a[:1]
	// ! mutates 'a' when this is enabled
	c := b[:2]
	// take from 0 idx to 1, len = 2, cap = 2
	// ! does not mutate 'a'
	// create a slice with length and capacity 2 only
	// c := b[0:2:2]

	fmt.Printf("a[%p] = %v\n", &a, a) // if array then can't reuse variable like %[1]v since Go doesn't allow it
	fmt.Printf("b[%p] = %[1]v\n", b)  // for slices param re-usage (%[1]v) is permitted
	fmt.Printf("c[%p] = %[1]v\n", c)
	fmt.Println("len(c) =", len(c))
	fmt.Println("cap(c) =", cap(c))
	fmt.Println("------------------------------")

	c = append(c, 5)
	fmt.Printf("mutated a[%p] = %v\n", &a, a) // a is mutated
	fmt.Printf("c[%p] = %[1]v\n", c)
	fmt.Println("------------------------------")

	c = append(c, 6)
	fmt.Println("new len(c) =", len(c))
	fmt.Println("new cap(c) =", cap(c))
	fmt.Printf("new c[%p] = %[1]v\n", c) // since the capacity of a is exceeded, a new memory is re-allocated for slice 'c' and the descriptor points to this address
	fmt.Println("------------------------------")

	c[0] = 7
	fmt.Println("len(c) =", len(c))
	fmt.Println("cap(c) =", cap(c))
	fmt.Printf("c[%p] = %[1]v\n", c)
	fmt.Println("------------------------------")
}
