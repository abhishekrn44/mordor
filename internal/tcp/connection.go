package tcp

import (
	"log"
	"net"
	"rana/mordor/http"
	"rana/mordor/parser"
	"rana/mordor/routes"
	"strings"
	"time"
)

func handleConnection(c net.Conn) {
	defer c.Close()

	log.Println("client connected", c.RemoteAddr())
	client_count++
	log.Println("active connections:", client_count)

	for {
		request, errCode := parser.ParseRequest(c)

		if errCode == -1 {
			log.Println("client closed connection")
			break
		}

		if errCode != 0 {
			if err := http.WriteResponse(http.NewErrorResponse(errCode, http.DeriveResponseHeaders(request)), c); err != nil {
				log.Println("error writing err:", err)
				break
			}
		}

		// based on request’s method, path, and target; return the resulting HTTP response.
		response := routes.Serve(request)
		c.SetWriteDeadline(time.Now().Add(10 * time.Second))

		if err := http.WriteResponse(response, c); err != nil {
			log.Println("response write error:", err)
		}

		connectionRequestHeader := strings.TrimSpace(request.Headers[http.HeaderConnection])
		connectionResponseHeader := strings.TrimSpace(response.Headers[http.HeaderConnection])

		if connectionRequestHeader == "close" {
			break
		}

		if connectionResponseHeader == "close" {
			break
		}

		c.SetReadDeadline(time.Now().Add(15 * time.Second))

	}
	client_count--
	log.Println("active connections:", client_count)
}
