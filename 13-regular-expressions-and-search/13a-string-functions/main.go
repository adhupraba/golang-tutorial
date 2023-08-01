package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	test := "Here is $1 which is $2!"

	test = strings.ReplaceAll(test, "$1", "honey")
	test = strings.ReplaceAll(test, "$2", "tasty")

	fmt.Println("replace all =>", test)
	fmt.Println("------------------------------------------")
	fmt.Println(A())
}

func A() string {
	return B()
}

func B() string {
	pc, file, line, ok := runtime.Caller(1)

	fmt.Println("program counter =>", pc)
	fmt.Println("file =>", file)
	fmt.Println("line =>", line)
	fmt.Println("ok =>", ok)

	idx := strings.LastIndexByte(file, '/')

	return "last index byte => " + file[idx+1:] + ":" + strconv.Itoa(line)
}
