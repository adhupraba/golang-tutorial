package main

import (
	"log"
	"time"
)

func main() {
	chans := []chan int{
		make(chan int),
		make(chan int),
	}

	for i := range chans {
		go func(i int, ch chan<- int) {
			for {
				time.Sleep(time.Duration(i) * time.Second)

				ch <- i
			}
		}(i+1, chans[i])
	}

	for i := 0; i < 12; i++ {

		// n0 := <-chans[0]
		// log.Println(i, "received n0", n0)

		// n1 := <-chans[1]
		// log.Println(i, "received n1", n1)

		select {
		case m0 := <-chans[0]:
			log.Println(i, "received m0", m0)
		case m1 := <-chans[1]:
			log.Println(i, "received m1", m1)
		}
	}
}
