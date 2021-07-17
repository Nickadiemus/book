package main

import (
	"book/methods/pointmethod/geometry"
	"fmt"
)

func main() {
	fmt.Println("point package")
	p := geometry.Point{5, 5}
	q := geometry.Point{2, 2}
	fmt.Println(p.Distance(q))           // same functionality, except this is package level (function call)
	fmt.Println(geometry.Distance(p, q)) // same functionality, except this type associative method (method call)
	w := geometry.Point{8, 6}
	pt := geometry.Path{p, q, w}
	fmt.Println(geometry.Path.Distance(pt))
}
