package http

import (
	"log"
	"strconv"
)

type Response struct {
	Version    string // e.g. HTTP/1.1
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Headers    map[string]string
	Body       []byte
}

// NewResponse builds a fully-formed HTTP response with sane defaults.
// - Sets Version to HTTP/1.1 if empty
// - Sets Status text from StatusCode (ReasonPhrase)
// - Ensures Headers map exists
// - Sets Content-Length based on body length
// - Sets Connection default to "close" if not provided
// - Merges extraHeaders (overrides defaults if same key)
func NewResponse(statusCode int, body []byte, extraHeaders map[string]string) *Response {
	if body == nil {
		body = []byte{}
	}

	headers := make(map[string]string, 8)

	// Defaults
	headers[HeaderContentLength] = strconv.Itoa(len(body))
	headers[HeaderConnection] = "close"

	// Merge caller headers (override defaults if needed)
	for k, v := range extraHeaders {
		headers[k] = v
	}

	return &Response{
		Version:    Version11,
		Status:     ReasonPhrase(statusCode),
		StatusCode: statusCode,
		Headers:    headers,
		Body:       body,
	}
}

// helper for error responses.
func NewErrorResponse(statusCode int, headers map[string]string) *Response {
	log.Println("creating error response for status code:", statusCode)
	body, ct := getBody(statusCode)

	if headers != nil {
		headers[HeaderContentType] = ct
	}

	return NewResponse(statusCode, body, headers)
}
