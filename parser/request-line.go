package parser

import (
	"fmt"
	"strings"
)

// Parses METHOD SP REQUEST-TARGET SP HTTP-VERSION CRLF
func ReadStartLine(data []byte) ([]string, int, error) {
	pos := 0
	var line strings.Builder

	for pos < len(data) && data[pos] != '\r' {
		line.WriteByte(data[pos])
		pos++
	}

	// Must end with CRLF
	if pos+1 >= len(data) || data[pos] != '\r' || data[pos+1] != '\n' {
		return nil, -1, fmt.Errorf("malformed request line (missing CRLF)")
	}

	startLine := line.String()
	parts := strings.Fields(startLine)

	if len(parts) != 3 {
		return nil, -1, fmt.Errorf("invalid request line: %q", startLine)
	}

	pos += 2 // consume \r\n
	return parts, pos, nil
}
