package main

import (
	"fmt"
	"sync"
)

// the goal is to print 1000
func do() int {
	var limit sync.Mutex
	var n int64
	var w sync.WaitGroup

	for i := 0; i < 1000; i++ {
		w.Add(1)

		go func() {
			limit.Lock()
			n++ // DATA RACE
			limit.Unlock()
			w.Done()
		}()
	}

	w.Wait()
	return int(n)
}

func main() {
	fmt.Println(do())
}
