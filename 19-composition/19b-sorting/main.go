package main

import (
	"fmt"
	"sort"
)

type Organ struct {
	Name   string
	Weight int
}

type Organs []Organ

type ByName struct{ Organs }

type ByWeight struct{ Organs }

func (s Organs) Len() int { return len(s) }

func (s Organs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByName) Less(i, j int) bool {
	return s.Organs[i].Name < s.Organs[j].Name
}

func (s ByWeight) Less(i, j int) bool {
	return s.Organs[i].Weight < s.Organs[j].Weight
}

func main() {
	s := []Organ{{"brain", 1340}, {"liver", 1494}, {"spleen", 162}, {"pancreas", 131}, {"heart", 290}}
	fmt.Println("original =>", s)

	sort.Sort(ByWeight{s})
	fmt.Println("sort by weight =>", s)

	sort.Sort(ByName{s})
	fmt.Println("sort by name =>", s)
}