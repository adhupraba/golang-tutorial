package main

import (
	"log"
	"time"
)

func main() {
	log.Println("start")

	tickRate := 2 * time.Second

	stopper := time.After(5 * tickRate)
	ticker := time.NewTicker(tickRate).C

	// we are naming the loop so that we can break from it
	// else break statement will break from the select only
loop:
	for {
		select {
		case <-ticker:
			log.Println("tick")

		case <-stopper:
			break loop
		}
	}

	log.Println("finish")
}
