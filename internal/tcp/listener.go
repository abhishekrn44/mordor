package tcp

import (
	"log"
	"net"
	"rana/mordor/config/server"
	"rana/mordor/http"
	"rana/mordor/parser"
	"strconv"
)

func StartHttpServer() {
	listener, err := net.Listen("tcp", server.Host+":"+strconv.Itoa(server.Port))

	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			panic(err)
		}

		log.Println("client connected", conn.RemoteAddr())

		response := parser.ProcessRequest(conn)

		err = http.WriteResponse(response, conn)

		// resp := "HTTP/1.1 200 OK\r\n" +
		// 	"Content-Type: text/html; charset=utf-8\r\n" +
		// 	"Content-Length: 236\r\n" +
		// 	"Connection: keep-alive\r\n" +
		// 	"\r\n" +
		// 	"<!DOCTYPE html><html><head><style>html { color-scheme: light dark; } body { width: 35em; margin: 0 auto;font-family: Tahoma, Verdana, Arial, sans-serif; }</style></head><body style = {background: black}><h1>Hello HTTP</h1></body></html>"

		// _, err = conn.Write([]byte(resp))

		if err != nil {
			log.Println(err)
		}

		log.Println("client disconnected", conn.RemoteAddr())
		conn.Close()
	}
}
