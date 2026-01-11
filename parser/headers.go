package parser

import (
	"fmt"
	"strings"
)

func ReadHeaders(data []byte, pos int) (map[string]string, int, error) {
	headers := make(map[string]string)

	for {
		// Need at least CRLF to proceed safely
		if pos >= len(data) {
			return nil, pos, fmt.Errorf("incomplete headers")
		}

		// Blank line = end of headers: \r\n
		if data[pos] == '\r' {
			if pos+1 >= len(data) || data[pos+1] != '\n' {
				return nil, pos, fmt.Errorf("malformed header terminator")
			}
			pos += 2 // consume \r\n
			break
		}

		var line strings.Builder
		for pos < len(data) && data[pos] != '\r' {
			line.WriteByte(data[pos])
			pos++
		}

		if pos+1 >= len(data) || data[pos] != '\r' || data[pos+1] != '\n' {
			return nil, pos, fmt.Errorf("malformed header line ending")
		}
		pos += 2 // consume \r\n

		headerLine := line.String()
		parts := strings.SplitN(headerLine, ":", 2)
		if len(parts) != 2 {
			return nil, pos, fmt.Errorf("invalid header line: %q", headerLine)
		}

		key := strings.ToLower(strings.TrimSpace(parts[0]))
		val := strings.TrimSpace(parts[1])

		// TODO: handle duplicate headers
		headers[key] = val

	}

	return headers, pos, nil
}
