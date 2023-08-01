package main

import (
	"fmt"
	"regexp"
	"strings"
)

var u = regexp.MustCompile(`^(http|https)://([a-zA-Z0-9\-\.]+\.[a-zA-Z]{2,4})(?::([0-9]+))?\/?([a-zA-Z0-9\-\._\?\,\'\/\\\+&%\$#\=~]*)$`)
var v = regexp.MustCompile(`(\w+)=(\w+)`)

var test = []string{
	"https://matt.com/hello",
	"http://matt.com:8080/hello/",
	"http://dom.dom.matt.com:8080/hello?a=1&b=2",
	"http:matt.com:8080/hello?a=1&b=2",
	"http://matt.com:8080/hello?a=1&b=2",
	"http://matt.com:8080/hello/?a=1&b=2&c=3",
}

func main() {
	// LiteralPrefix -> is there actually a prefix that has to be there all the time
	r, b := u.LiteralPrefix()

	fmt.Printf("%q %t\n", r, b)

	for i, t := range test {
		match := u.FindStringSubmatch(t)
		fmt.Printf("%d: %q\n", i, match)

		if len(match) > 4 && strings.Contains(match[4], "?") {
			vars := v.FindAllStringSubmatch(match[4], -1)

			fmt.Printf("   %q\n", vars)
		}
	}
}
