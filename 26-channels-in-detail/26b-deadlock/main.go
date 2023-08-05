package main

import "fmt"

func main() {
	// ! works
	// // buffered channel with capacity 2, can write two data
	// messages := make(chan string, 2)

	// // write 1st data
	// messages <- "hello"
	// // write 2nd data, success because there is enough capacity
	// messages <- "world"

	// // read 1st data, free up 1 space in channel
	// fmt.Println(<-messages)
	// // read 2nd data, free up all the space in channe;
	// fmt.Println(<-messages)

	// ! does not work
	// buffered channel with capacity 1, can write 1 data only
	// can write additional data only when already written data is read from the channel
	messages := make(chan string, 1)

	// write 1st data, takes up the available 1 space in channel
	messages <- "hello"
	// try to write 2nd data, not possible because channel capacity is 1 only
	// already 1st data is still there waiting to be read
	// when encountering the below write what happens is control flow gets blocked here and does not go through
	// the read from channel is after the 2nd write only. so there is a deadlock. the control flow cannot proceed
	// the program crashes
	messages <- "world"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
