package routes

import (
	"log"
	"os"
	"path/filepath"
	"rana/mordor/http"
	"strconv"
	"strings"
)

func Serve(req *http.Request) *http.Response {
	path := req.Target

	if path == "/" {
		path = "static/default.html"
	}

	// Prevent path traversal
	if strings.Contains(path, "..") {
		return http.ErrorResponse(http.StatusNotFound)
	}

	data, err := os.ReadFile(path)

	if err != nil {
		log.Println("read file error:", err)
		return http.ErrorResponse(http.StatusNotFound)
	}

	ct := detectContentType(path)

	return &http.Response{
		StatusCode: 200,
		Status:     "OK",
		Headers: map[string]string{
			"Content-Type":   ct,
			"Content-Length": strconv.Itoa(len(data)),
			"Connection":     "close",
		},
		Body: data,
	}
}

func detectContentType(path string) string {
	switch filepath.Ext(path) {
	case ".html":
		return http.ContentTypeTextPlain
	case ".css":
		return http.ContentTypeCSS
	case ".js":
		return http.ContentTypeJS
	case ".png":
		return http.ContentTypePNG
	case ".jpg", ".jpeg":
		return http.ContentTypeJPEG
	default:
		return http.ContentTypeOctet
	}
}
