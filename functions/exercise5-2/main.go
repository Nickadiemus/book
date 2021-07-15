package main

import (
	"fmt"
	"os"

	"book/functions/exercise5-2/count"

	"golang.org/x/net/html"
)

func isError(err error) bool {
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		return true
	}
	return false
}

func main() {
	f, err := os.Open("./golang.html")

	if isError(err) {
		os.Exit(1)
	}
	doc, derr := html.Parse(f)
	if isError(derr) {
		os.Exit(1)
	}

	fmt.Printf("<html>: %d\n", count.CountElements(0, "html", doc))
	fmt.Printf("<head>: %d\n", count.CountElements(0, "head", doc))
	fmt.Printf("<div>: %d\n", count.CountElements(0, "div", doc))
	fmt.Printf("<span>: %d\n", count.CountElements(0, "span", doc))
	fmt.Printf("<p>: %d\n", count.CountElements(0, "p", doc))
	fmt.Printf("<ul>: %d\n", count.CountElements(0, "ul", doc))
	fmt.Printf("<li>: %d\n", count.CountElements(0, "li", doc))
	fmt.Printf("<a>: %d\n", count.CountElements(0, "a", doc))

}
