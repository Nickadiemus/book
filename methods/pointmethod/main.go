package main

import (
	"book/methods/pointmethod/geometry"
	"fmt"
)

func main() {
	fmt.Println("point package")
	p := geometry.Point{5, 5}
	q := geometry.Point{2, 2}
	fmt.Println(p.Distance(q))
}
