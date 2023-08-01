package main

import (
	"fmt"
	"os"

	hello "2-simple-example/helper"
)

// go run ./2-simple-example adharsh
func main() {
	if len(os.Args) > 1 {
		fmt.Println(hello.SayHello(os.Args[1:]))
	} else {
		names := []string{"world"}
		fmt.Println(hello.SayHello(names))
	}
}
