package http

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

func WriteError(response *Response, w io.Writer) error {

	bw := bufio.NewWriter(w)

	if response.StatusCode != 0 && response.Status != "" {
		if _, err := fmt.Fprintf(bw, "%s %d %s\r\n", response.Version, response.StatusCode, response.Status); err != nil {
			return err
		}
	} else {
		if _, err := fmt.Fprintf(bw, "%s %s\r\n", response.Version, strings.TrimSpace(response.Status)); err != nil {
			return err
		}
	}

	// Headers (sorted for stable output)
	keys := make([]string, 0, len(response.Headers))
	for k := range response.Headers {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := response.Headers[k]
		if _, err := fmt.Fprintf(bw, "%s: %s\r\n", k, v); err != nil {
			return err
		}
	}

	// End headers
	if _, err := bw.WriteString("\r\n"); err != nil {
		return err
	}

	// Body
	if len(response.Body) > 0 {
		if _, err := bw.Write(response.Body); err != nil {
			return err
		}
	}

	bw.Available()

	return bw.Flush()
}
