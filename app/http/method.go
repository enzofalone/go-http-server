package http

const (
	MethodGET    = "GET"
	MethodPOST   = "POST"
	MethodUPDATE = "UPDATE"
	MethodDELETE = "DELETE"
)

type PathMethods struct {
	GET    func() string
	POST   func() string
	DELETE func() string
	UPDATE func() string
}
