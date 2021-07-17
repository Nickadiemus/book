package main

import (
	"book/programstructure/exercise2-1/tempconv"
	"flag"
	"fmt"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(temp)
}
