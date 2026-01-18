package http

import (
	"os"
	"rana/mordor/config/server"
)

var body404 []byte
var body500 []byte

func init() {
	body404, _ = os.ReadFile(server.BaseDir + "static/404.html")
	body500, _ = os.ReadFile(server.BaseDir + "static/500.html")
}

func getBody(code int) ([]byte, string) {

	switch code {
	case 404:
		return body404, ContentTypeHTML
	case 500:
		return body500, ContentTypeHTML
	default:
		return []byte(ReasonPhrase(code)), ContentTypeTextPlain
	}

}
