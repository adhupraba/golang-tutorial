package main

import "fmt"

func main() {
	ch := make(chan int, 1) // buffered channel
	// ch := make(chan string, 1)

	// ! below unbuffered channel will crash because the second read (c, ok := <-ch) is blocked
	// ! since it is waiting for a write to happen so that it can read
	// ! but the channel is closed and no one can write into it, so it is deadlock
	// ch := make(chan int)

	ch <- 1
	// ch <- "apple"

	b, ok := <-ch
	fmt.Println(b, ok) // 1 true

	close(ch)

	// ! an unbuffered channel requires a reader and writer (a writer blocked on a channel with no reader will "leak")

	// we can read from a closed channel
	// when reading from a closed channel the value returned is the default value for the channel's data type
	// eg: for int - default value = 0, string - default value = "", etc..
	c, ok := <-ch
	fmt.Println(c, ok) // 0 false
}
