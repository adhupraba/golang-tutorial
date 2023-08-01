package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IntSlice []int

type Sample struct {
	x int
}

// ! methods can be put only on user declared types, eg: IntSlice
// ! methods can also be put on structs. but in this example we are not doing it, so error will be thrown
// ! methods cannot be put on built in types directly, eg: int, string

func (is IntSlice) String() string {
	var strs []string

	for _, v := range is {
		strs = append(strs, strconv.Itoa(v))
	}

	return "[" + strings.Join(strs, ";") + "]"
}

func main() {
	var v IntSlice = []int{1, 2, 3}

	// * works because IntSlice type has String() method which we declared above
	var s fmt.Stringer = v

	// var sample = Sample { x: 10 }
	// ! error becase Sample struct does not not contain the String() method which is required for fmt.Stringer interface
	// var s fmt.Stringer = sample

	// ! error becase []int type does not not contain the String() method which is required for fmt.Stringer interface
	// var s fmt.Stringer = []int{1, 2, 3}

	for i, x := range v {
		fmt.Printf("%d: %d\n", i, x)
	}

	fmt.Printf("%T %[1]v\n", v)
	fmt.Printf("%T %[1]v\n", s)
}
