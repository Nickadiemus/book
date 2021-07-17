package main

import (
	"book/interfaces/httphandler/database"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Initializing....")
	db := database.Database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.List)
	http.HandleFunc("/price", db.Price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
