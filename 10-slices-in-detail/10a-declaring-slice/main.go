package main

import "fmt"

func main() {
	var s []int
	t := []int{}
	u := make([]int, 5)
	v := make([]int, 0, 5)

	fmt.Printf("s -> %d, %d, %T, %t, %#[3]v\n", len(s), cap(s), s, s == nil)
	fmt.Printf("t -> %d, %d, %T, %t, %#[3]v\n", len(t), cap(t), t, t == nil)
	fmt.Printf("u -> %d, %d, %T, %t, %#[3]v\n", len(u), cap(u), u, u == nil)
	fmt.Printf("v -> %d, %d, %T, %t, %#[3]v\n", len(v), cap(v), v, v == nil)
}
