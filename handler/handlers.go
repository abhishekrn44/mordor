package handler

import (
	"log"
	"net/http"
	"os"
)

var welcome string = "static/default.html"
var notfound string = "static/404.html"

func HandleRequest(path string) ([]byte, int, int) {

	if path == home {
		return readResource(welcome)
	}

	return readResource(notfound)
}

func readResource(path string) ([]byte, int, int) {
	body, err := os.ReadFile(path)

	if err != nil {
		log.Println("resource read error:", err)
		return nil, -1, http.StatusInternalServerError
	}
	return body, len(body), http.StatusOK
}
