package main

import (
	"book/goroutines/concurrentecho/server/echo"
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	fmt.Println("Starting Echo Function...")
	// create listen
	listener, err := net.Listen("tcp", "localhost:8000")
	// check if initialization error
	if err != nil {
		log.Fatal(err)
	}
	// while
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g client aborted
			continue
		}
		handleConn(conn)
	}
	// use listener to check for client connection
	// if EOF (^C) is sent from client terminate
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)

	for input.Scan() {
		go echo.Echo(c, input.Text(), 1*time.Second)
	}
	defer c.Close()
}
