package main

import "fmt"

// ! THIS EXAMPLE IS PRETTY CONFUSING
// ! GO THROUGH THE OUTPUT IN TERMINAL TO UNDERSTAND THE FLOW

// ch chan <- int    ----- write into channel
// src <- chan int   ----- read from channel

func generator(limit int, ch chan<- int) {
	for i := 2; i < limit; i++ {
		fmt.Printf("generator i = %v, ch = %#v\n", i, ch)
		ch <- i
	}

	fmt.Println("------- closing generator channel -------")

	close(ch)
}

func filter(src <-chan int, dst chan<- int, prime int) {
	fmt.Printf("enter filter src chan = %#v, dst = %v, prime = %v\n", src, dst, prime)

	for i := range src {
		fmt.Printf("filter i = %v, prime = %v, src = %#v, dst = %#v\n", i, prime, src, dst)

		if i%prime != 0 {
			fmt.Printf("######### src = %#v, dst = %#v, prime = %v, pushing i = %v into dst channel #########\n", src, dst, prime, i)
			dst <- i
		}
	}

	fmt.Printf("------- closing filter channel <<>> src = %#v, dst = %#v, prime = %v -------\n", src, dst, prime)

	close(dst)
}

func sieve(limit int) {
	src := make(chan int)
	fmt.Printf("sieve initial src = %#v\n", src)

	go generator(limit, src)

	for {
		// prime - value, ok - boolean (is the channel closed or not)
		prime, ok := <-src

		fmt.Printf("sieve read from src channel.. src = %#v, prime = %v, ok = %v\n", src, prime, ok)

		if !ok {
			fmt.Printf("<<<<<< breaking for loop src = %#v, prime = %v >>>>>>\n", src, prime)
			break
		}

		dst := make(chan int)

		fmt.Printf("sieve for loop src = %#v, dst = %#v, prime = %v\n", src, dst, prime)

		go filter(src, dst, prime)

		src = dst

		fmt.Printf(".......... prime result => %v ..........\n", prime)
	}
}

func main() {
	sieve(100) // 2 3 5 7 11 13 17 19 23 ...
}
