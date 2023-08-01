package main

import "fmt"

func do(d func()) {
	d()
}

func main() {
	for i := 0; i < 4; i++ {
		v := func() {
			fmt.Printf("%d @ %p\n", i, &i)
		}

		do(v)
	}

	fmt.Println("------------------------")

	s1 := make([]func(), 4)

	for i := 0; i < 4; i++ {
		s1[i] = func() {
			fmt.Printf("%d @ %p\n", i, &i)
		}
	}

	for i := 0; i < 4; i++ {
		s1[i]()
	}

	fmt.Println("------------------------")

	s2 := make([]func(), 4)

	for i := 0; i < 4; i++ {
		j := i // closure capture
		s2[i] = func() {
			fmt.Printf("%d @ %p\n", j, &j)
		}
	}

	for i := 0; i < 4; i++ {
		s2[i]()
	}
}
