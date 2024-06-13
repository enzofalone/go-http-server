package http

const (
	StatusOK                  = 200 // RFC 9110, 15.3.1
	StatusAccepted            = 202 // RFC 9110, 15.3.3
	StatusBadRequest          = 400 // RFC 9110, 15.5.1
	StatusUnauthorized        = 401 // RFC 9110, 15.5.2
	StatusForbidden           = 403 // RFC 9110, 15.5.4
	StatusNotFound            = 404 // RFC 9110, 15.5.5
	StatusMethodNotAllowed    = 405 // RFC 9110, 15.5.6
	StatusInternalServerError = 500 // RFC 9110, 15.6.1
)

func StatusText(code int) string {
	switch code {
	case StatusOK:
		return "OK"
	case StatusAccepted:
		return "Accepted"
	case StatusBadRequest:
		return "Bad Request"
	case StatusUnauthorized:
		return "Unauthorized"
	case StatusForbidden:
		return "Forbidden"
	case StatusNotFound:
		return "Not Found"
	case StatusMethodNotAllowed:
		return "Method Not Allowed"
	case StatusInternalServerError:
		return "Internal Server Error"
	default:
		return ""
	}
}
