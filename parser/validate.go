package parser

import (
	"rana/mordor/http"
	"slices"
)

// Validate request for supported HTTP Version and Method.
func ValidateStartLine(method, version string) int {

	if version != http.Version11 {
		return http.StatusHTTPVersionNotSupported
	}

	if !slices.Contains(http.Methods, method) {
		return http.StatusNotImplemented
	}

	return 0

}
