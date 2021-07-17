package main

import (
	"book/goroutines/channels/counter/counter"
	"book/goroutines/channels/counter/squarer"
	"fmt"
)

func main() {
	fmt.Println("Initializing....")
	natural := make(chan int)
	squares := make(chan int)
	go counter.Counter(natural)
	go squarer.Squarer(squares, natural)
	printer(squares)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
