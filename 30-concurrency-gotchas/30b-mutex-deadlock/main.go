package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Mutex
	done := make(chan bool)

	fmt.Println("START")

	go func() {
		m.Lock()
		defer m.Unlock() // if commented, deadlock occurs
	}()

	go func() {
		time.Sleep(1 * time.Nanosecond)

		m.Lock()
		defer m.Unlock()

		fmt.Println("SIGNAL")
		done <- true
	}()

	<-done
	fmt.Println("DONE")
}
