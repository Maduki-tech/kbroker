package main

import (
	"log"
	"net"

	"github.com/maduki-tech/kbroker/internal/server"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(server.Serve(listener))
}
