package main

import "strings"

func mapQueryParams(queryParamString string) map[string]string {
	queryParamsArr := strings.Split(queryParamString, "&")

	m := make(map[string]string)

	for _, queryParam := range queryParamsArr {
		slice := strings.Split(queryParam, "=")

		k := slice[0]
		v := slice[1]

		m[k] = v
	}

	return m
}

// utility function which extracts params in route according to ":" prefix
func mapParams(path string, route string) map[string]string {
	params := make(map[string]string)

	pathParts := strings.Split(path, "/")
	routeParts := strings.Split(route, "/")

	if len(pathParts) != len(routeParts) {
		return nil
	}

	for i, routePart := range routeParts {
		// found param
		if strings.HasPrefix(routePart, ":") {
			paramName := routePart[1:]
			paramValue := pathParts[i]

			params[paramName] = paramValue
		} else if routeParts[i] != pathParts[i] {
			// it does not match
			return nil
		}
	}

	return params
}
