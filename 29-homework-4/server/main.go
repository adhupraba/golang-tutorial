package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type dollars float32

type database struct {
	mu sync.Mutex
	db map[string]dollars
}

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

// add the handlers
func (d *database) list(w http.ResponseWriter, r *http.Request) {
	d.mu.Lock()
	defer d.mu.Unlock()

	for item, price := range d.db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (d *database) getItem(w http.ResponseWriter, r *http.Request) {
	d.mu.Lock()
	defer d.mu.Unlock()

	item := r.URL.Query().Get("item")

	if _, ok := d.db[item]; !ok {
		msg := fmt.Sprintf("no such item: %q", item)
		http.Error(w, msg, http.StatusNotFound) // 404
		return
	}

	fmt.Fprintf(w, "item %s with price %s\n", item, d.db[item])
}

func (d *database) add(w http.ResponseWriter, r *http.Request) {
	d.mu.Lock()
	defer d.mu.Unlock()

	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")

	if _, ok := d.db[item]; ok {
		msg := fmt.Sprintf("duplicate item: %q", item)
		http.Error(w, msg, http.StatusBadRequest) // 400
		return
	}

	p, err := strconv.ParseFloat(price, 32)

	if err != nil {
		msg := fmt.Sprintf("invalid price: %q", price)
		http.Error(w, msg, http.StatusBadRequest) // 400
		return
	}

	d.db[item] = dollars(p)

	fmt.Fprintf(w, "added %s with price %s\n", item, d.db[item])
}

func (d *database) update(w http.ResponseWriter, r *http.Request) {
	d.mu.Lock()
	defer d.mu.Unlock()

	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")

	if _, ok := d.db[item]; !ok {
		msg := fmt.Sprintf("no such item: %q", item)
		http.Error(w, msg, http.StatusNotFound) // 404
		return
	}

	p, err := strconv.ParseFloat(price, 32)

	if err != nil {
		msg := fmt.Sprintf("invalid price: %q", price)
		http.Error(w, msg, http.StatusBadRequest) // 400
		return
	}

	d.db[item] = dollars(p)

	fmt.Fprintf(w, "updated %s with price %s\n", item, d.db[item])
}

func (d *database) delete(w http.ResponseWriter, r *http.Request) {
	d.mu.Lock()
	defer d.mu.Unlock()

	item := r.URL.Query().Get("item")

	if _, ok := d.db[item]; !ok {
		msg := fmt.Sprintf("no such item: %q", item)
		http.Error(w, msg, http.StatusNotFound) // 404
		return
	}

	delete(d.db, item)

	fmt.Fprintf(w, "item %s deleted\n", item)
}

func main() {
	db := database{
		db: map[string]dollars{
			"shoes": 50,
			"socks": 5,
		},
	}

	// add some routes
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/read", db.getItem)
	http.HandleFunc("/create", db.add)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)

	fmt.Println("server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
