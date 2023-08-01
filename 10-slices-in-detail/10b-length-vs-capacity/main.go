package main

import "fmt"

func main() {
	// a := []int{1, 2, 3} // slice
	// a := [...]int{1} // array, program will crash if this is enabled
	a := [...]int{1, 2, 3} // array
	fmt.Println("a =", a)
	fmt.Println("len(a) =", len(a))
	fmt.Println("cap(a) =", cap(a))

	fmt.Println("--------------------")

	b := a[:1]
	fmt.Println("b =", b)
	fmt.Println("len(b) =", len(b))
	fmt.Println("cap(b) =", cap(b))

	fmt.Println("--------------------")

	// since b is created from a and b has a descriptor which points to the same slice as a
	// c is able to point to first two elements of a through b
	// ! b and c pick up the underlying capacity of a
	c := b[:2]
	fmt.Println("c =", c)
	fmt.Println("len(c) =", len(c))
	fmt.Println("cap(c) =", cap(c))

	fmt.Println("--------------------")

	d := c[0:1:1] // [i:j:k] -> i = start index to grab from the parent array, len = j-i, cap = k-i
	// ! if capacity is 0, then even if u give any index, slice will not be populated with value
	fmt.Println("d =", d)
	fmt.Println("len(d) =", len(d))
	fmt.Println("cap(d) =", cap(d))

	fmt.Println("--------------------")

	// e := d[0:2] // d only has capacity of 1, so slicing 2 elements out of it will throw error
	e := d[0:1]
	fmt.Println("e =", e)
	fmt.Println("len(e) =", len(e))
	fmt.Println("cap(e) =", cap(e))

	fmt.Println("--------------------")

}
