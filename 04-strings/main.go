package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// search and replace program
// go run ./main.go matt ed < test.txt
// go run ./main.go matt ed (and type sentences in command line.. ctrl+d to quit the program)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "not enough args")
		os.Exit(1)
	}

	// os.Args[0] is the name of the program in the cli
	old, new := os.Args[1], os.Args[2]

	scan := bufio.NewScanner(os.Stdin)

	// scan.Scan() -> returns true if there is a line present
	for scan.Scan() {
		// scan.Text() -> returns the text itself
		s := strings.Split(scan.Text(), old)
		t := strings.Join(s, new)

		fmt.Println(t)
	}
}
