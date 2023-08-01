package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	te := "aba abba abbba"
	re := regexp.MustCompile(`b+`)
	mm := re.FindAllString(te, -1)
	id := re.FindAllStringIndex(te, -1)

	fmt.Println("found sub strings =>", mm)
	fmt.Println("found sub string indices", id)

	fmt.Println("-------- sub strings ---------")

	for _, d := range id {
		fmt.Println(te[d[0]:d[1]])
	}

	fmt.Println("------------------------------")

	up := re.ReplaceAllStringFunc(te, strings.ToUpper)

	fmt.Println("new to upper =>", up)
}
