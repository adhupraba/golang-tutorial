package main

import (
	"fmt"
	"os"
)

// to read input from nums.txt, run
// go run ./main.go < nums.txt
// cat nums.txt | go run ./main.go
func main() {
	var sum float64
	var n int

	fmt.Println("enter the numbers. to calculate the average and exit the program press `ctrl+d`")

	for {
		var val float64

		_, err := fmt.Fscanln(os.Stdin, &val)

		if err != nil {
			fmt.Printf("error is %v\n", err)
			break
		}

		sum += val
		n++
	}

	if n == 0 {
		fmt.Fprintln(os.Stderr, "no values")
		os.Exit(1)
	}

	fmt.Println("the average is", sum/float64(n))
}
