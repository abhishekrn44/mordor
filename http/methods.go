package http

const (
	MethodGet  = "GET"
	MethodPost = "POST"
	MethodPut  = "PUT"
	MethodDel  = "DELETE"
	MethodHead = "HEAD"
	MethodOpts = "OPTIONS"
)

var Methods []string = []string{MethodGet, MethodPost, MethodPut, MethodDel, MethodHead, MethodOpts}
