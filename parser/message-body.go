package parser

import (
	"bufio"
	"io"
	"log"
	"rana/mordor/http"
	"strconv"
)

func readBody(r *bufio.Reader, headers map[string]string) ([]byte, int) {
	cl, ok := headers["content-length"]
	if !ok {
		return nil, 0
	}

	n, err := strconv.Atoi(cl)
	if err != nil || n < 0 {
		log.Println("invalid content-length:", cl)
		return nil, http.StatusBadRequest
	}

	if n == 0 {
		return nil, 0
	}

	body := make([]byte, n)
	if _, err := io.ReadFull(r, body); err != nil {
		log.Println("read body error:", err)
		return nil, http.StatusBadRequest
	}
	return body, 0
}
