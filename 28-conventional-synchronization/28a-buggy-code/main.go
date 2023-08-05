package main

import (
	"fmt"
	"sync"
)

// the goal is to print 1000
func do() int {
	var n int64
	var w sync.WaitGroup

	for i := 0; i < 1000; i++ {
		w.Add(1)

		go func() {
			// since all the goroutines are running concurrently there is huge chance that
			// some goroutines can access the same value stored in the variable and increment at the same time
			// and some increments will be lost
			// eg: 3 goroutines access 'n' which has a value of 800
			// ideally post incremenet we expect 803 to be present but the value would have been incremented to 801 only
			// because all 3 goroutines had 800 as the value of 'n' when 'n' was accessed to increment
			// every goroutine incremented from 800 to 801
			n++ // DATA RACE
			w.Done()
		}()
	}

	w.Wait()
	return int(n)
}

func main() {
	fmt.Println(do())
}
