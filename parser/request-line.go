package parser

import (
	"fmt"
	"log"
	"net"
	"rana/mordor/http"
	"strings"
)

func ReadMessage(conn net.Conn) (http.Request, error) {

	var buf []byte = make([]byte, 1024)

	count, err := conn.Read(buf[:])

	if err != nil {
		log.Fatalln(err)
		return http.Request{}, err
	}

	var str strings.Builder
	pos := 0

	for ; pos < count; pos++ {
		if buf[pos] == '\r' {
			break
		}
		str.WriteByte(buf[pos])
	}

	startLine := str.String()
	comps := strings.Split(startLine, " ")

	if len(comps) < 3 {
		return http.Request{}, fmt.Errorf("invalid request line")
	}

	pos += 2
	str.Reset()

	eof := false
	headers := make(map[string]string)

	for !eof {
		eof, err, pos = readHeaders(buf, pos, headers)
		if err != nil {
			return http.Request{}, fmt.Errorf("invalid request line")
		}
		pos += 2
	}

	request := http.Request{
		Method:      comps[0],
		Target:      comps[1],
		Version:     comps[2],
		Headers:     headers,
		MessageBody: nil,
	}

	return request, nil
}

func readHeaders(buf []byte, pos int, headers map[string]string) (bool, error, int) {

	if buf[pos] == '\r' {
		return true, nil, pos
	}

	var str strings.Builder

	for ; pos < len(buf); pos++ {
		if buf[pos] == '\r' {
			break
		}
		str.WriteByte(buf[pos])
	}
	header := str.String()
	header_comps := strings.SplitN(header, ":", 2)

	if len(header_comps) != 2 {
		return false, fmt.Errorf("invalid header line"), pos
	}

	headers[header_comps[0]] = header_comps[1]
	return false, nil, pos
}
