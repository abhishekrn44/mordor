package http

const (
	StatusOK                  = 200
	StatusBadRequest          = 400
	StatusNotFound            = 404
	StatusMethodNotAllowed    = 405
	StatusRequestTimeout      = 408
	StatusPayloadTooLarge     = 413
	StatusInternalServerError = 500
	StatusNotImplemented      = 501
)

var StatusText = map[int]string{
	StatusOK:                  "OK",
	StatusBadRequest:          "Bad Request",
	StatusNotFound:            "Not Found",
	StatusMethodNotAllowed:    "Method Not Allowed",
	StatusRequestTimeout:      "Request Timeout",
	StatusPayloadTooLarge:     "Payload Too Large",
	StatusInternalServerError: "Internal Server Error",
	StatusNotImplemented:      "Not Implemented",
}

func ReasonPhrase(code int) string {
	if s, ok := StatusText[code]; ok {
		return s
	}
	return ""
}
