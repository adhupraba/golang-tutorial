package main

import "fmt"

func main() {
	// a, b := 12, 345
	// c, d := 1.2, 3.45

	// fmt.Printf("%d %d\n", a, b)
	// base16 or hexadecimal
	// fmt.Printf("%x %x\n", a, b)
	// fmt.Printf("%X %X\n", a, b)
	// fmt.Printf("%#x %#x\n", a, b)
	// fmt.Printf("%f %.2f\n", c, d)

	// -----------------------------

	// fmt.Printf("|%6d|%6d|\n", a, b)
	// fmt.Printf("|%06d|%06d|\n", a, b)
	// fmt.Printf("|%-6d|%-6d|\n", a, b)
	// fmt.Printf("|%6f|%6.2f|\n", c, d)

	// -----------------------------

	slice := []int{1, 2, 3}
	array := [3]rune{'a', 'b', 'c'}
	mapp := map[string]int{"and": 1, "or": 2}
	str := "a string"
	bytes := []byte(str)

	fmt.Println("----- slice ------")
	fmt.Printf("%T\n", slice)
	fmt.Printf("%v\n", slice)
	fmt.Printf("%#v\n", slice)

	fmt.Println("----- array ------")
	fmt.Printf("%T\n", array)
	fmt.Printf("%q\n", array)
	fmt.Printf("%v\n", array)
	fmt.Printf("%#v\n", array)

	fmt.Println("----- map ------")
	fmt.Printf("%T\n", mapp)
	fmt.Printf("%v\n", mapp)
	fmt.Printf("%#v\n", mapp)

	fmt.Println("----- string ------")
	fmt.Printf("%T\n", str)
	fmt.Printf("%q\n", str)
	fmt.Printf("%v\n", str)
	fmt.Printf("%#v\n", str)

	fmt.Println("----- bytes ------")
	fmt.Printf("%T\n", bytes)
	fmt.Printf("%q\n", bytes)
	fmt.Printf("%v\n", bytes)
	fmt.Printf("%#v\n", bytes)
	fmt.Printf("bytes -> string: %v\n", string(bytes))
}
