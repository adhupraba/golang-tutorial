package main

import (
	"fmt"
	"sync"
)

// the goal is to print 1000
func do() int {
	// counting semaphore which limits active goroutines to 1
	limit := make(chan bool, 1)
	var n int64
	var w sync.WaitGroup

	for i := 0; i < 1000; i++ {
		w.Add(1)

		go func() {
			limit <- true
			n++ // DATA RACE
			<-limit
			w.Done()
		}()
	}

	w.Wait()
	return int(n)
}

func main() {
	fmt.Println(do())
}
