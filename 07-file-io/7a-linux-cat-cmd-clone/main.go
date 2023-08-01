package main

import (
	"fmt"
	"io"
	"os"
)

// go run main.go a.txt b.txt c.txt
// go run main.go *.txt
// go run main.go *.txt > d.txt
// ! be sure to delete d.txt before running the above command

func main() {
	fmt.Println("args =>", os.Args[1:])

	for _, fname := range os.Args[1:] {
		fmt.Println("filename", fname)
		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		if _, err := io.Copy(os.Stdout, file); err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}

		file.Close()
	}
}
