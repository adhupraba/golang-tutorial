package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// go run main.go a.txt b.txt c.txt
// go run main.go *.txt

func main() {
	for _, fname := range os.Args[1:] {
		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		data, err := ioutil.ReadAll(file)

		if err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}

		fmt.Println("the file", fname, "has", len(data), "bytes")

		file.Close()
	}
}
