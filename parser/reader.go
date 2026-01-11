package parser

import (
	"bufio"
	"io"
	"log"
	"net"
	"rana/mordor/http"
	"strconv"
	"strings"
)

func ReadRequest(conn net.Conn) (*http.Request, *http.Response) {
	r := bufio.NewReader(conn)

	// Read start-line
	startLine, err := r.ReadString('\n')
	if err != nil {
		log.Println("read start-line error:", err)
		status := http.ReasonPhrase(http.StatusBadRequest)
		len := strconv.Itoa(len(status))
		response := &http.Response{
			Version:    http.Version11,
			StatusCode: http.StatusBadRequest,
			Status:     status,
			Headers: map[string]string{
				http.HeaderContentType:   http.ContentTypeTextPlain,
				http.HeaderConnection:    "close",
				http.HeaderContentLength: len,
			},
			Body: []byte(status),
		}
		return &http.Request{}, response
	}
	log.Println("Start-Line:", strings.TrimRight(startLine, "\r\n"))

	startLineParts := strings.Split(startLine, " ")

	// Read headers
	headers := make(map[string]string)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			log.Println("read header error:", err)
			status := http.ReasonPhrase(http.StatusBadRequest)
			len := strconv.Itoa(len(status))
			response := &http.Response{
				Version:    http.Version11,
				StatusCode: http.StatusBadRequest,
				Status:     status,
				Headers: map[string]string{
					http.HeaderContentType:   http.ContentTypeTextPlain,
					http.HeaderConnection:    "close",
					http.HeaderContentLength: len,
				},
				Body: []byte(status),
			}
			return &http.Request{}, response
		}

		if line == "\r\n" { // blank line = end of headers
			break
		}

		// Parse: Header - Key: Value
		raw := strings.TrimRight(line, "\r\n")
		kv := strings.SplitN(raw, ":", 2)
		if len(kv) != 2 {
			status := http.ReasonPhrase(http.StatusBadRequest)
			len := strconv.Itoa(len(status))
			response := &http.Response{
				Version:    http.Version11,
				StatusCode: http.StatusBadRequest,
				Status:     status,
				Headers: map[string]string{
					http.HeaderContentType:   http.ContentTypeTextPlain,
					http.HeaderConnection:    "close",
					http.HeaderContentLength: len,
				},
				Body: []byte(status),
			}
			return &http.Request{}, response
		}

		key := strings.ToLower(strings.TrimSpace(kv[0]))
		val := strings.TrimSpace(kv[1])

		// TODO: handle duplicates
		headers[key] = val
	}

	// Read body if present
	body := []byte{}
	if cl, ok := headers["content-length"]; ok {
		n, err := strconv.Atoi(cl)
		if err != nil || n < 0 {
			log.Println("invalid Content-Length")
			status := http.ReasonPhrase(http.StatusBadRequest)
			len := strconv.Itoa(len(status))
			response := &http.Response{
				Version:    http.Version11,
				StatusCode: http.StatusBadRequest,
				Status:     status,
				Headers: map[string]string{
					http.HeaderContentType:   http.ContentTypeTextPlain,
					http.HeaderConnection:    "close",
					http.HeaderContentLength: len,
				},
				Body: []byte(status),
			}
			return &http.Request{}, response
		}
		body = make([]byte, n)
		_, err = io.ReadFull(r, body)
		if err != nil {
			log.Println("read body error:", err)
			status := http.ReasonPhrase(http.StatusBadRequest)
			len := strconv.Itoa(len(status))
			response := &http.Response{
				Version:    http.Version11,
				StatusCode: http.StatusBadRequest,
				Status:     status,
				Headers: map[string]string{
					http.HeaderContentType:   http.ContentTypeTextPlain,
					http.HeaderConnection:    "close",
					http.HeaderContentLength: len,
				},
				Body: []byte(status),
			}
			return &http.Request{}, response
		}
	}

	if len(body) > 0 {
		log.Println("Body:", string(body))
	}

	request := &http.Request{
		Method:      startLineParts[0],
		Target:      startLineParts[1],
		Version:     startLineParts[2],
		Headers:     headers,
		MessageBody: body,
	}

	if request.Target != "/" {
		status := http.ReasonPhrase(http.StatusBadRequest)
		len := strconv.Itoa(len(status))
		response := &http.Response{
			Version:    http.Version11,
			StatusCode: http.StatusBadRequest,
			Status:     status,
			Headers: map[string]string{
				http.HeaderContentType:   http.ContentTypeTextPlain,
				http.HeaderConnection:    "close",
				http.HeaderContentLength: len,
			},
			Body: []byte(status),
		}
		return &http.Request{}, response
	}

	return request, &http.Response{}
}
