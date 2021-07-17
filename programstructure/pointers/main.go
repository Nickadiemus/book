package main

import "fmt"

func main() {

	x := 1
	p := &x                    // p, of type *int, points to x
	fmt.Println("x:", x)       // "1"
	fmt.Println("p:", p)       // "p contains the address of x"
	fmt.Println("*p:", *p)     // "1"
	*p = 2                     // equivalent to x = 2
	fmt.Println("after x:", x) // "2"
}
