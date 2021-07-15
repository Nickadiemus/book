package main

import (
	"fmt"
	"os"

	"book/functions/exercise5-1/visit"

	"golang.org/x/net/html"
)

func main() {
	data, ferr := os.Open("./golang.html")
	if ferr != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", ferr)
		os.Exit(1)
	}
	doc, err := html.Parse(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(visit.VisitR(nil, doc))
}
