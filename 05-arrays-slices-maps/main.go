package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// search and replace program
// go run ./main.go matt ed < test.txt
// go run ./main.go matt ed (and type sentences in command line.. ctrl+d to quit the program)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)

	scan.Split(bufio.ScanWords)

	// scan.Scan() -> returns true if there is a line present
	for scan.Scan() {
		// scan.Text() -> returns the text itself
		// here we get a word because we are splitting the input on the words
		// see line number 18
		// fmt.Println("word =>", scan.Text())
		words[scan.Text()]++
	}

	// fmt.Println(words)
	fmt.Println(len(words), "unique words")

	type kv struct {
		key string
		val int
	}

	var ss []kv

	for k, v := range words {
		// ss = append(ss, kv{key: k, val: v})
		// ss = append(ss, kv{val: v, key: k})
		ss = append(ss, kv{k, v})
	}

	fmt.Println("ss =>", ss)

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].val > ss[j].val
	})

	for _, s := range ss[:3] {
		fmt.Println(s.key, "appears", s.val, "times")
	}
}
