package http

type Response struct {
	Version    string // e.g. HTTP/1.1
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Headers    map[string]string
	Body       []byte
}
