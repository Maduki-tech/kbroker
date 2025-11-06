// Package server
package server

import (
	"errors"
	"io"
	"net"

	"github.com/maduki-tech/kbroker/internal/logstore"
	"github.com/maduki-tech/kbroker/internal/protocol"
)

func Serve(listener net.Listener) error {
	for {
		conn, err := listener.Accept()
		if err != nil {
			if _, ok := err.(net.Error); ok {
				continue
			}
			return err
		}

		go handleConnection(conn)

	}
}

func handleConnection(conn net.Conn) {
	packet := make([]byte, 4096)
	for {
		n, err := conn.Read(packet)
		if err != nil {
			if err == io.EOF {
				return
			}

			var ne net.Error
			if errors.As(err, &ne) {
				return
			}

			return
		}

		if n == 0 {
			return
		}

		msg, err := protocol.ParseMessage(string(packet[:n]))
		if err != nil {
			// Skip to not drop connection on parse error
			continue
		}

		logstore.Write(msg)

	}
}
