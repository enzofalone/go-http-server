package main

import (
	"log"
	"strings"

	"github.com/enzofalone/go-http-server/app/handlers"
	"github.com/enzofalone/go-http-server/app/http"
)

type handlerMap map[string]map[string]func(http.Request) string

// function that finds and executes handler matched by pathname
func findHandler(path string, method string) string {
	// iterate over every route
	for route, methodHandlers := range handleFuncs {
		// trim query params
		splitPath := strings.Split(path, "?")

		var queryParams map[string]string
		if len(splitPath) != 1 {
			queryParams = mapQueryParams(splitPath[1])
		}

		params := mapParams(splitPath[0], route)

		req := http.Request{QueryParams: queryParams, Params: params}

		// execute if:
		// path is exact
		// path extracts params from route
		if path == route || params != nil {
			methodHandler, ok := methodHandlers[method]

			if !ok {
				return handlers.NotFound(http.Request{})
			}

			return methodHandler(req)
		}
	}
	return handlers.NotFound(http.Request{})
}

// function to add handler into handlerFuncs with error checking
// TODO: maybe add tests???
// TODO: i believe it is important for handlers that have path parameters
// to have some sort of collision check within the handleFuncs maps
func (hmap handlerMap) addHandler(method string, path string, handleFunc handlers.Handler) {
	if method != http.MethodGET &&
		method != http.MethodPOST &&
		method != http.MethodUPDATE &&
		method != http.MethodDELETE {
		log.Fatal("method " + method + " does not exist!")
	}

	// initialize inner method map for every new path
	if handleFuncs[path] == nil {
		handleFuncs[path] = make(map[string]func(http.Request) string)
	}

	// handle duplicate paths
	if _, ok := handleFuncs[path][method]; ok {
		log.Fatal("path \"" + path + "\" already exists!")
	}

	handleFuncs[path][method] = handleFunc
}

func setupHandlers(hmap handlerMap) {
	hmap.addHandler(http.MethodGET, "/", handlers.GenericError)
	hmap.addHandler(http.MethodGET, "/echo/:str", handlers.HandleEcho)
	hmap.addHandler(http.MethodGET, "/ping", handlers.HandleHealth)
}
