package http

import (
	"strconv"
)

func ErrorResponse(code int) *Response {

	switch code {
	case 400:
		return badRequestResponse()
	case 404:
		return notFoundResponse()
	case 500:
		return internalServerErrorResponse()
	case 501:
		return notImplementedResponse()
	case 505:
		return notImplementedResponse()
	default:
		return nil
	}

}

func badRequestResponse() *Response {
	status := ReasonPhrase(StatusBadRequest)
	contentLen := strconv.Itoa(len(status))

	return &Response{
		Version:    Version11,
		StatusCode: StatusBadRequest,
		Status:     status,
		Headers: map[string]string{
			HeaderContentType:   ContentTypeTextPlain,
			HeaderConnection:    "close",
			HeaderContentLength: contentLen,
		},
		Body: []byte(status),
	}
}

func notImplementedResponse() *Response {
	status := ReasonPhrase(StatusNotImplemented)
	contentLen := strconv.Itoa(len(status))

	return &Response{
		Version:    Version11,
		StatusCode: StatusNotImplemented,
		Status:     status,
		Headers: map[string]string{
			HeaderContentType:   ContentTypeTextPlain,
			HeaderConnection:    "close",
			HeaderContentLength: contentLen,
		},
		Body: []byte(status),
	}
}

func httpVersionNotSupportedResponse() *Response {
	status := ReasonPhrase(StatusHTTPVersionNotSupported)
	contentLen := strconv.Itoa(len(status))

	return &Response{
		Version:    Version11,
		StatusCode: StatusHTTPVersionNotSupported,
		Status:     status,
		Headers: map[string]string{
			HeaderContentType:   ContentTypeTextPlain,
			HeaderConnection:    "close",
			HeaderContentLength: contentLen,
		},
		Body: []byte(status),
	}
}

func internalServerErrorResponse() *Response {
	status := ReasonPhrase(StatusInternalServerError)
	contentLen := strconv.Itoa(len(status))

	return &Response{
		Version:    Version11,
		StatusCode: StatusInternalServerError,
		Status:     status,
		Headers: map[string]string{
			HeaderContentType:   ContentTypeTextPlain,
			HeaderConnection:    "close",
			HeaderContentLength: contentLen,
		},
		Body: []byte(status),
	}

}

func notFoundResponse() *Response {
	status := ReasonPhrase(StatusNotFound)
	contentLen := strconv.Itoa(len(status))

	return &Response{
		Version:    Version11,
		StatusCode: StatusNotFound,
		Status:     status,
		Headers: map[string]string{
			HeaderContentType:   ContentTypeTextPlain,
			HeaderConnection:    "close",
			HeaderContentLength: contentLen,
		},
		Body: []byte(status),
	}
}
