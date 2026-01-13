package parser

import (
	"bufio"
	"io"
	"log"
	"rana/mordor/http"
	"strconv"
)

func readBody(r *bufio.Reader, headers map[string]string) ([]byte, *http.Response) {
	cl, ok := headers["content-length"]
	if !ok {
		return nil, nil
	}

	n, err := strconv.Atoi(cl)
	if err != nil || n < 0 {
		log.Println("invalid content-length:", cl)
		return nil, http.ErrorResponse(http.StatusBadRequest)
	}

	if n == 0 {
		return nil, nil
	}

	body := make([]byte, n)
	if _, err := io.ReadFull(r, body); err != nil {
		log.Println("read body error:", err)
		return nil, http.ErrorResponse(http.StatusBadRequest)
	}
	return body, nil
}
