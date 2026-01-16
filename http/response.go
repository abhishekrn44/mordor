package http

import "strconv"

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
	headers[HeaderConnection] = "close"
	headers[HeaderContentLength] = strconv.Itoa(len(body))

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

// helper for common text responses.
func NewTextResponse(statusCode int, text string) *Response {
	return NewResponse(statusCode, []byte(text), map[string]string{
		HeaderContentType: ContentTypeTextPlain,
	})
}

// returns a minimal text/plain error response.
// If msg is empty, uses the standard reason phrase as the body.
func NewErrorResponse(statusCode int, msg string) *Response {
	if msg == "" {
		msg = ReasonPhrase(statusCode)
	}
	return NewResponse(statusCode, []byte(msg), map[string]string{
		HeaderContentType: ContentTypeTextPlain,
	})
}
