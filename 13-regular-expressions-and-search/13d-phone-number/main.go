package main

import (
	"fmt"
	"regexp"
)

// american phone number pattern
var pattern = `\(([[:digit:]]{3})\) ([[:digit:]]{3})-([[:digit:]]{3})`
var phre = regexp.MustCompile(pattern)

func main() {
	// LiteralPrefix -> is there actually a prefix that has to be there all the time
	// b = false -> this is not the entire string
	// b = true -> this is the entire string
	r, b := phre.LiteralPrefix()

	fmt.Printf("%q %t\n", r, b)

	original := "(214) 514-9548"
	match := phre.FindStringSubmatch(original)

	fmt.Printf("%q\n", match)

	// print local phone number to international phone number format
	if len(match) > 3 {
		fmt.Printf("+1 %s-%s-%s\n", match[1], match[2], match[3])
	}

	fmt.Println("------------------------------------------")

	orig := "call me at (214) 514-9548 today"
	intl := phre.ReplaceAllString(orig, "+1 ${1}-${2}-${3}")

	fmt.Println(intl)
}
