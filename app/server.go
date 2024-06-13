package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/enzofalone/go-http-server/app/handlers"
	"github.com/enzofalone/go-http-server/app/http"
)

// handlers will be stored in memory and initialized at startup
// they return a string which is basically the http response
// handleFuncs contains two hashmaps, one hashmap for path
// second hashmap for methods supported by that path
// as a result, it returns a function
// eg: handleFuncs["/user"]["GET"] -> GetUserHandler()
var handleFuncs = make(map[string]map[string]func() string)

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

	// check if there are path parameters
	// do this

	response := execHandler(method, path)

	// if path == "/" {
	// 	response = "HTTP/1.1 200 OK\r\n\r\n"
	// } delete this

	if _, err := conn.Write([]byte(response)); err != nil {
		log.Fatal(err)
	}
}

func findHandler(method string, path string) bool {
	_, ok := handleFuncs[path][method]
	return ok
}

// based on method and path of request, execute handler previously set up in setupHandlers function
func execHandler(method string, path string) string {
	handlerFunc, ok := handleFuncs[path][method]

	if ok {
		return handlerFunc()
	}

	// path does not exist
	return handlers.NotFound()
}

// function to add handler into handlerFuncs with error checking
// TODO: maybe add tests???
func addHandler(method string, path string, handleFunc func() string) {
	if method != http.MethodGET &&
		method != http.MethodPOST &&
		method != http.MethodUPDATE &&
		method != http.MethodDELETE {
		log.Fatal("method " + method + " does not exist!")
	}

	// initialize inner method map for every new path
	if handleFuncs[path] == nil {
		handleFuncs[path] = make(map[string]func() string)
	}

	// handle duplicate paths
	if _, ok := handleFuncs[path][method]; ok {
		log.Fatal("path \"" + path + "\" already exists!")
	}

	handleFuncs[path][method] = handleFunc
}

func setupHandlers() {
	addHandler(http.MethodGET, "/ping", handlers.HandleHealth)
}

func main() {
	setupHandlers()

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
