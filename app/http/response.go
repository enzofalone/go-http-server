package http

import (
	"fmt"
	"strconv"
)

// construct HTTP/1.1 response based on data and status code
func RespondWithStatus(data string, code int) string {
	response := "HTTP/1.1 " + strconv.Itoa(code) + " " + StatusText(code) + "\r\n"

	// header stuff
	// maybe add some state variable in the future for more personalization
	response += "Content-Type: text/plain\r\n"
	response += "Content-Length: " + strconv.Itoa(len(data)) + "\r\n\r\n"

	//response body
	if len(data) > 0 {
		response += data
	}
	fmt.Println(response)
	return response
}
