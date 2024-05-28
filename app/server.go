package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func handleConnection(conn net.Conn) {
	// free up memory by closing connection once satisfied
	defer conn.Close()
	buf := make([]byte, 1024)
	conn.Read(buf)

	bufString := strings.Split(string(buf), "\n")

	req := strings.Split(bufString[0], " ")
	// headers := strings.Split(bufString[1], " ")

	// method := req[0]
	path := req[1]
	// version := req[2]
	fmt.Printf("Path is %s\n", path)

	var response string = "HTTP/1.1 404 Not Found\r\n\r\n"
	if path == "/" {
		response = "HTTP/1.1 200 OK\r\n\r\n"
	}

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
