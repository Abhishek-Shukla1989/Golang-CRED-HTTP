package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollor float64

func (d dollor) String() string {

	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollor

func (db database) list(w http.ResponseWriter, r *http.Request) {

	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) delete(w http.ResponseWriter, r *http.Request) {

	item := r.URL.Query().Get("item")

	if _, ok := db[item]; !ok {

		msg := fmt.Sprintf("Item didn't exist %q", item)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	delete(db, item)
	fmt.Fprintf(w, "Item deleted %q ", item)
}

func (db database) add(w http.ResponseWriter, r *http.Request) {

	// get params from query first
	// check if item exists in db, send error
	// add item to db
	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")

	if _, ok := db[item]; ok {

		msg := fmt.Sprintf("Duplicate item %q", item)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		msg := fmt.Sprintf("Invalid price %q %q", item, price)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	db[item] = dollor(p)
	fmt.Fprintf(w, "Item added %q with price %q", item, price)

}

func (db database) update(w http.ResponseWriter, r *http.Request) {

	// get params from query first
	// check if item exists in db, send error
	// add item to db
	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")

	if _, ok := db[item]; !ok {

		msg := fmt.Sprintf("Ietm didn't exist %q", item)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		msg := fmt.Sprintf("Invalid price %q %q", item, price)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	db[item] = dollor(p)
	fmt.Fprintf(w, "Item updated %q with price %q", item, price)

}

func main() {

	db := database{
		"shoe":  50,
		"socks": 100,
	}

	http.HandleFunc("/list", db.list)
	http.HandleFunc("/create", db.add)
	http.HandleFunc("/delete", db.delete)
	http.HandleFunc("/update", db.update)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
