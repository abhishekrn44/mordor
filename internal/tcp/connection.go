package tcp

import (
	"log"
	"net"
	"rana/mordor/http"
	"rana/mordor/parser"
	"strings"
)

func handleConnection(c net.Conn) {
	defer c.Close()

	log.Println("client connected", c.RemoteAddr())

	for {
		request, errCode := parser.ProcessRequest(c)
		_ = request

		if errCode != 0 {
			if err := http.WriteResponse(http.ErrorResponse(errCode), c); err != nil {
				log.Println("error writing err:", err)
				break
			}
		}

		// TODO: select the appropriate handler based on the request’s method and target, invokes it, and return the resulting HTTP response.
		// response := handler.HandleRequest(request)

		var response *http.Response

		connectionRequestHeader := strings.TrimSpace(request.Headers[http.HeaderConnection])
		connectionResponseHeader := strings.TrimSpace(response.Headers[http.HeaderConnection])

		if connectionRequestHeader == "close" {
			break
		}

		if connectionResponseHeader == "close" {
			break
		}

	}

	log.Println("client disconnected", c.RemoteAddr())
}
