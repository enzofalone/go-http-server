package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

// handlers will be stored in memory and initialized at startup
// they return a string which is basically the http response
// handleFuncs contains two hashmaps, one hashmap for path
// second hashmap for methods supported by that path
// as a result, it returns a function
// eg: handleFuncs["/user"]["GET"] -> GetUserHandler()
var handleFuncs = make(handlerMap)

// main function that dissects request into variables
// also finds which handler to use based on path/request
func resolveConnection(conn net.Conn) {
	// free up memory by closing connection once satisfied
	defer conn.Close()

	buf := make([]byte, 1024)
	conn.Read(buf)

	bufString := strings.Split(string(buf), "\n")

	req := strings.Split(bufString[0], " ")
	// headers := strings.Split(bufString[1], " ")

	method := req[0]
	path := req[1]

	response := findHandler(path, method)

	if _, err := conn.Write([]byte(response)); err != nil {
		log.Fatal(err)
	}
}

func main() {
	setupHandlers(handleFuncs)

	port := 4221
	l, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer l.Close()

	fmt.Println("Running on port " + strconv.Itoa(port))

	for {
		// wait for any new connections
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		// handle new connection by passing all information
		go resolveConnection(conn)
	}

}
