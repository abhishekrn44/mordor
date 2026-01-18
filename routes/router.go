package routes

import (
	"log"
	"os"
	"path/filepath"
	"rana/mordor/config/server"
	"rana/mordor/http"
	"strings"
)

func Serve(req *http.Request) *http.Response {
	path := req.Target

	log.Println("Serving for path:", path)

	// TODO: route handling to be added here
	if path == "/" {
		path = "static/default.html"
	}

	if path == "/emp" {
		path = "emp/index.html"
	}

	// Prevent path traversal
	if strings.Contains(path, "..") {
		return http.NewErrorResponse(http.StatusNotFound, http.DeriveResponseHeaders(req))
	}

	fullPath := filepath.Join(server.BaseDir, filepath.Clean(path))

	data, err := os.ReadFile(fullPath)

	if err != nil {
		log.Println("read file error:", err)
		return http.NewErrorResponse(http.StatusNotFound, http.DeriveResponseHeaders(req))
	}

	ct := detectContentType(path)

	return http.NewResponse(http.StatusOK, data, map[string]string{
		http.HeaderContentType: ct,
		http.HeaderConnection:  req.Headers[strings.ToLower(http.HeaderConnection)],
	})

}

func detectContentType(path string) string {
	switch filepath.Ext(path) {
	case ".html":
		return http.ContentTypeHTML
	case ".css":
		return http.ContentTypeCSS
	case ".js":
		return http.ContentTypeJS
	case ".png":
		return http.ContentTypePNG
	case ".jpg", ".jpeg":
		return http.ContentTypeJPEG
	case ".svg":
		return http.ContentTypeSVG
	default:
		return http.ContentTypeOctet
	}
}
