package main

import (
	"fmt"
	"log"
	"net/http"
)

type nextCh chan int

func (ch nextCh) handler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("handler func", req.URL.Path)
	fmt.Fprintf(w, "<h1>You got %d</h1>", <-ch)
}

// ch chan <- int    ----- write into channel
// src <- chan int   ----- read from channel
func counter(ch chan<- int) {
	for i := 0; ; i++ {
		fmt.Println("writing to channel", i)
		ch <- i
	}
}

func main() {
	var nextID nextCh = make(chan int)

	go counter(nextID)

	fmt.Println("inside main func")

	http.HandleFunc("/", nextID.handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
