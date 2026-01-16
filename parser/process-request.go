package parser

import (
	"bufio"
	"log"
	"net"
	"rana/mordor/http"
)

func ProcessRequest(conn net.Conn) (*http.Request, int) {

	log.Println("Processing request")

	r := bufio.NewReader(conn)

	method, target, version, code := readStartLine(r)

	log.Printf("%s %s %s", method, target, version)

	if code != 0 {
		return nil, code
	}

	if code := ValidateStartLine(method, version); code != 0 {
		return nil, code
	}

	headers, code := readHeaders(r)

	if code != 0 {
		return nil, code
	}

	messageBody, code := readBody(r, headers)

	if code != 0 {
		return nil, code
	}

	request := &http.Request{
		Method:      method,
		Target:      target,
		Version:     version,
		Headers:     headers,
		MessageBody: messageBody,
	}

	log.Println("Reqest", request)

	return request, 0

}
