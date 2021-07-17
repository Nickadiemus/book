package main

import (
	"book/interfaces/bytecounter/bcount"
	"fmt"
)

func main() {
	var c bcount.ByteCounter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}
