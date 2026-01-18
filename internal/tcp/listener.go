package tcp

import (
	"log"
	"net"
	"rana/mordor/config/server"
	"strconv"
)

var client_count int = 0
var max_clients int = 20000

func StartHttpServer() {
	listener, err := net.Listen("tcp", server.Host+":"+strconv.Itoa(server.Port))

	if err != nil {
		panic(err)
	}

	for {

		if client_count >= max_clients {
			log.Println("Maximun clients limit hit! current active client:", client_count)
			continue
		}

		conn, err := listener.Accept()

		if err != nil {
			panic(err)
		}

		go handleConnection(conn)
	}

}
