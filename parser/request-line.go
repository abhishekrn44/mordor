package parser

import (
	"bufio"
	"io"
	"log"
	"net"
	"rana/mordor/http"
	"strings"
)

// func ReadStartLine(data []byte) ([]string, int, error) {
// 	pos := 0
// 	var line strings.Builder

// 	for pos < len(data) && data[pos] != '\r' {
// 		line.WriteByte(data[pos])
// 		pos++
// 	}

// 	// Must end with CRLF
// 	if pos+1 >= len(data) || data[pos] != '\r' || data[pos+1] != '\n' {
// 		return nil, -1, fmt.Errorf("malformed request line (missing CRLF)")
// 	}

// 	startLine := line.String()
// 	parts := strings.Fields(startLine)

// 	if len(parts) != 3 {
// 		return nil, -1, fmt.Errorf("invalid request line: %q", startLine)
// 	}

// 	pos += 2 // consume \r\n
// 	return parts, pos, nil
// }

// Parses METHOD SP REQUEST-TARGET SP HTTP-VERSION CRLF
func readStartLine(r *bufio.Reader) (method, target, version string, statusCode int) {
	log.Println("Reading start-line")
	line, err := r.ReadString('\n')
	if err != nil {
		log.Println("read start-line error:", err)

		if err == io.EOF {
			return "", "", "", -1
		} else if ne, ok := err.(net.Error); ok && ne.Timeout() {
			return "", "", "", -1
		} else {
			return "", "", "", http.StatusBadRequest
		}
	}

	line = strings.TrimRight(line, "\r\n")
	parts := strings.Fields(line)
	if len(parts) != 3 {
		log.Println("invalid request line:", line)
		return "", "", "", http.StatusBadRequest
	}

	return parts[0], parts[1], parts[2], 0
}
