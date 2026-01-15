package tcp

import (
	"net"
	"rana/mordor/config/server"
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

		go handleConnection(conn)

	}
}
