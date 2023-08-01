package main

import "fmt"

type errFoo struct {
	err  error
	path string
}

func (e errFoo) Error() string {
	return fmt.Sprintf("%s: %s", e.path, e.err)
}

// func XYZ(a int) *errFoo {
// 	return nil
// }

func XYZ(a int) error {
	return nil
}

func main() {
	// err := XYZ(1) // err would be *errFoo, not error interface
	var err error = XYZ(1) // BAD: interface gets a nil concrete pointer
	fmt.Println(err)

	if err != nil {
		fmt.Println("oops")
	} else {
		fmt.Println("ok")
	}
}
