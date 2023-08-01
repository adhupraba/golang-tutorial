package main

import (
	"fmt"
	"regexp"
)

// var uu = regexp.MustCompile(`[[:xdigit:]]{8}-[[:xdigit:]]{4}-[1-5][[:xdigit:]]{3}-[89abAB][[:xdigit:]]{3}-[[:xdigit:]]{12}`)
var uu = regexp.MustCompile(`^[[:xdigit:]]{8}-[[:xdigit:]]{4}-[1-5][[:xdigit:]]{3}-[89abAB][[:xdigit:]]{3}-[[:xdigit:]]{12}$`)

var test = []string{
	"072665ee-a034-4cc3-a2e8-9f1822c4ebbb",
	"072665ee-a034-6cc3-a2e8-9f1822c4ebbb",  // invalid version
	"072665ee-a034-4cc3-72e8-9f1822c4ebbb",  // invalid type
	"072665ee-a034-4cc3-a2e8-9f1822c4ebb",   // less character
	"072665ee-a034-4cc3-a2e8-9f1822c4ebbcb", // extra character
	"072665ee-a034-3cc3-82e8-9f1822c4ebbb",
}

func main() {
	for i, t := range test {
		if !uu.MatchString(t) {
			fmt.Println(i, t, "\tfails")
		}
	}
}
