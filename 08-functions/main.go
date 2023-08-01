package main

import "fmt"

/*
array -> pass by value
slice -> pass by reference (actually the slice descriptor is copied and sent as parameter, not the slice itself as reference)
map -> (m1, m2)two map descriptors pointing to the same hash table. when a change is made using m2 variable
			m1 map gets altered. but if a new map is initialized into the m2 map then the descriptor which points
			to m1 map now points to a new map in memory. the changes made previously to the m1 map stays and further
			changes from the point of new initialization affects only the new m2 map
			! kind of like pass by reference
*/

// b -> formal parameter
func doArr(b [3]int) int {
	b[0] = 0
	return b[1]
}

func doSlice(b []int) int {
	b[0] = 0
	// ! %p is used to print pointer
	fmt.Printf("slice b@ %p\n", b)
	return b[1]
}

func doMap(m1 map[int]int) {
	m1[3] = 0
	m1 = make(map[int]int)
	m1[4] = 4
	fmt.Println("m1", m1)
}

// this is explicit pass by reference
func doMapPointer(m1 *map[int]int) {
	(*m1)[3] = 0
	*m1 = make(map[int]int)
	(*m1)[4] = 4
	fmt.Println("m1", *m1)
}

func main() {
	arr := [3]int{1, 2, 3}
	// arr -> actual parameter
	v1 := doArr(arr)

	slice := []int{1, 2, 3}
	fmt.Printf("slice in main @ %p\n", slice)
	v2 := doSlice(slice)

	mapp := map[int]int{4: 1, 7: 2, 8: 3}
	doMap(mapp)

	mapPointer := map[int]int{4: 1, 7: 2, 8: 3}
	fmt.Println("map pointer before", mapPointer)
	doMapPointer(&mapPointer)

	fmt.Println("array", arr, v1)
	fmt.Println("slice", slice, v2)
	fmt.Println("mapp", mapp)
	fmt.Println("map pointer after", mapPointer)
}
