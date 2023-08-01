package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// go run main.go a.txt b.txt c.txt
// go run main.go *.txt

func main() {
	for _, fname := range os.Args[1:] {
		// line count, word count, character count
		var lc, wc, cc int

		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		scan := bufio.NewScanner(file)

		for scan.Scan() {
			s := scan.Text()
			// fmt.Println("value of s is =>", s)

			wc += len(strings.Fields(s))
			cc += len(s)
			lc++
		}

		fmt.Printf("%7d %7d %7d %7s\n", lc, wc, cc, fname)

		file.Close()
	}
}
