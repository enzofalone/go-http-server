package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	// free up memory by closing connection once satisfied
	defer conn.Close()

	response := "HTTP/1.1 200 OK\r\n\r\n"

	if _, err := conn.Write([]byte(response)); err != nil {
		log.Fatal(err)
	}

}

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", ":4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer l.Close()

	for {
		// wait for any new connections
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		// handle new connection by passing all information
		go handleConnection(conn)
	}

}
