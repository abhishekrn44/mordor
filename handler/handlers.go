package handler

import (
	"log"
	"net/http"
	"os"
)

func HandleRequest(path string) ([]byte, int, int) {

	if path == home {
		body, err := os.ReadFile("static/default.html")

		if err != nil {
			log.Println("resource read error:", err)
			return nil, -1, http.StatusInternalServerError
		}

		return body, len(body), http.StatusOK
	}

	return nil, -1, http.StatusNotFound
}
