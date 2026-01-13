package parser

import (
	"bufio"
	"log"
	"net"
	"rana/mordor/handler"
	"rana/mordor/http"
	"strconv"
)

func ProcessRequest(conn net.Conn) *http.Response {

	r := bufio.NewReader(conn)

	method, target, version, response := readStartLine(r)

	log.Printf("%s %s %s", method, target, version)

	if response != nil {
		return response
	}

	if code := ValidateStartLine(method, version); code != 0 {
		return http.ErrorResponse(code)
	}

	content, size, errCode := handler.HandleRequest(target)

	if content == nil {
		return http.ErrorResponse(http.StatusNotFound)
	}

	if errCode != http.StatusOK {
		return http.ErrorResponse(errCode)
	}

	headers, response := readHeaders(r)

	if response != nil {
		return response
	}

	messageBody, response := readBody(r, headers)
	_ = messageBody

	if response != nil {
		return response
	}

	status := http.ReasonPhrase(http.StatusOK)

	return &http.Response{
		Version:    http.Version11,
		StatusCode: http.StatusOK,
		Status:     status,
		Headers: map[string]string{
			http.HeaderContentType:   http.ContentTypeHTML,
			http.HeaderConnection:    "close",
			http.HeaderContentLength: strconv.Itoa(size),
		},
		Body: content,
	}

}
