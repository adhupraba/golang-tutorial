package main

import (
	"fmt"
	"path/filepath"
)

type Pair struct {
	Path string
	Hash string
}

type PairWithLength struct {
	Pair
	Length int
}

type FileNamer interface {
	Filenameee() string
}

// struct embedding pointer to another struct
// promotion of its fields works the same way
type Fizgig struct {
	*PairWithLength
	Broken bool
}

func (p Pair) String() string {
	return fmt.Sprintf("Hash of %s is %s", p.Path, p.Hash)
}

// if this function did not exist fmt.Println will use the String method of Pair type
func (p PairWithLength) String() string {
	return fmt.Sprintf("Hash of %s is %s; length %d", p.Path, p.Hash, p.Length)
}

func FileName(p Pair) string {
	return filepath.Base(p.Path)
}

func (p Pair) Filenameee() string {
	return filepath.Base(p.Path)
}

func main() {
	p := Pair{"/usr", "0xfdfe"}
	pl := PairWithLength{Pair{"/usr/lib", "0xdead"}, 10}
	// this assignment worked because Pair was promoted in PairWithLength type and Pair has the Filenameee function
	var fn FileNamer = PairWithLength{Pair{"/var/git", "0xtsbo"}, 21}
	fg := Fizgig{
		&PairWithLength{Pair{"/etc/nginx", "0xksjce"}, 100},
		false,
	}

	fmt.Println(p)
	fmt.Println(pl)
	fmt.Println("------------------")

	fmt.Println(FileName(p))
	// ! this will not work because FileName() takes a parameter with the exact type Pair
	// ! unlike interface methods which requires any type of object to contain that single method
	// fmt.Println(FileName(pl))
	// ! this will work
	fmt.Println(FileName(pl.Pair))
	fmt.Println("------------------")

	fmt.Println(p.Filenameee())
	fmt.Println(fn)
	fmt.Println(fn.Filenameee())
	fmt.Println("------------------")

	fmt.Println(fg)
}
