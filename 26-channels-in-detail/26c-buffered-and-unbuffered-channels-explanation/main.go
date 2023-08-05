package main

import (
	"fmt"
	"time"
)

type T struct {
	i byte
	b bool
}

func send(i int, ch chan<- *T) {
	t := &T{i: byte(i)}
	// if unbuffered channel, then write is blocked until the read is completed, so making change to 't' will not reflect
	// if buffered channel, then write is non-blocking, any change made to 't' is reflected
	ch <- t

	// in unbufered channel, the value is already read from the channel before reaching here, so any change made here will not reflect after read
	// in buffered channel, there is not blocking. after writing into channel 't' will be modified immediately and when reading the change is reflected
	t.b = true
}

func main() {
	vs := make([]T, 5)
	// ch := make(chan *T) // unbuffered channel
	ch := make(chan *T, 5) // buffered channel

	for i := range vs {
		go send(i, ch)
	}

	time.Sleep(1 * time.Second) // sleeping so all goroutines get started

	// copy quickly
	for i := range vs {
		vs[i] = *<-ch
	}

	// print
	for _, v := range vs {
		fmt.Println(v)
	}
}
