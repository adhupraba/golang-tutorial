package main

import (
	"fmt"
	"log"
	"net/http"
)

var nextID = make(chan int)

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("handler func", req.URL.Path)
	fmt.Fprintf(w, "<h1>You got %d</h1>", <-nextID)
}

// this will not immediately generate infinite numbers and push to the channel
// data can be written into a channel only when there is a reader (line 13)
// only when nextID is read when requesting the html, data will be written into the channel from counter function
// now the write part is on hold until a next read is made
// refer - https://youtu.be/zJd7Dvg3XCk?t=1371
func counter() {
	// all the iteratations of the loop is not done immediately
	// the control flow is paused until a data is read from the channel
	// only after data is read (line 13), the control flow resumes execution (next iteration starts)
	for i := 0; ; i++ {
		fmt.Println("writing to channel", i)
		nextID <- i
	}
}

func main() {
	go counter()
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
