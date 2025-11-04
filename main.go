package main

import (
	"net"

	"github.com/maduki-tech/kbroker/internal/logstore"
	"github.com/maduki-tech/kbroker/internal/protocol"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go handleConnection(conn)

	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	packet := make([]byte, 4096)
	for {
		n, err := conn.Read(packet)
		if err != nil {
			panic(err)
		}
		msg, err := protocol.ParseMessage(string(packet[:n]))
		if err != nil {
			panic(err)
		}

		logstore.Write(msg)

	}
}
