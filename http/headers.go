package http

import "strings"

const (
	HeaderHost          = "Host"
	HeaderContentType   = "Content-Type"
	HeaderContentLength = "Content-Length"
	HeaderConnection    = "Connection"
	HeaderAllow         = "Allow"
)

func DeriveResponseHeaders(request *Request) map[string]string {

	h := make(map[string]string)

	// Connection handling
	if v, ok := request.Headers["connection"]; ok && strings.ToLower(v) == "close" {
		h["Connection"] = "close"
	} else {
		h["Connection"] = "keep-alive"
	}

	return h

}
