package database

import (
	"fmt"
	"net/http"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type Database map[string]dollars

func (db Database) List(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db Database) Price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // this writes the standard 404 error code
		// since Fprintf takes any implementation of a .Writer method, we can just use
		// httpResonseWrite since it has a .Write() method
		fmt.Fprintf(w, "No such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
