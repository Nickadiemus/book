package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	fmt.Println("Starting Sequential Clock....")
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
		go handleConn(conn)
	}
	// use listener to check for client connection
	// if EOF (^C) is sent from client terminate
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
