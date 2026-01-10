package tcp

import (
	"log"
	"net"
	"rana/mordor/config/server"
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

		_, err = parser.ReadMessage(conn)
		if err != nil {
			conn.Close()
			log.Fatalln(err)
			continue
		}

		response := "HTTP/1.1 200 OK\r\n" +
			"Content-Type: text/html; charset=utf-8\r\n" +
			"Content-Length: 45\r\n" +
			"Connection: keep-alive\r\n" +
			"\r\n" +
			"<html><body><h1>Hello HTTP</h1></body></html>"

		_, err = conn.Write([]byte(response))

		if err != nil {
			log.Fatalln(err)
		}

		conn.Close()
		log.Println("client disconnected", conn.RemoteAddr())
	}
}
