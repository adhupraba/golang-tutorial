package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("sleeping")
	time.Sleep(7 * time.Second)
	log.Println("sending response")

	fmt.Fprintf(w, "<h1>This is localhost</h1>")
}

func main() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
