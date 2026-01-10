package main

import (
	"fmt"
	"log"
	"rana/mordor/config/server"
	"rana/mordor/internal/tcp"
	"strconv"
)

func main() {
	printBanner()
	log.Println("[MORDOR] One Server to rule them all now accepting requests.")
	tcp.StartHttpServer()
}

func printBanner() {
	fmt.Print(`
============================================================
               M O R D O R
             HTTP Server Rises
============================================================
Listening on:  http://` + server.Host + `:` + strconv.Itoa(server.Port) + `
Realm:         Middle-Earth Networking
============================================================

`)
}
