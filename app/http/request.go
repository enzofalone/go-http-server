package http

// in Request struct, all http standard values will be stored here
type Request struct {
	QueryParams map[string]string
	Params      map[string]string
}
