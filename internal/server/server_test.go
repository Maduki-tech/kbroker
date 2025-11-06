// Package server_test
package server_test

import (
	"net"
	"testing"
	"time"

	"github.com/maduki-tech/kbroker/internal/server"
)

// // If your wire format needs something else, adjust payload accordingly.
// // Keep it small to reduce kernel buffer backpressure during the bench.
var payload = []byte("PRODUCE Hello")

func BenchmarkEndToEndProduce(b *testing.B) {
	ln, _ := net.Listen("tcp", "127.0.0.1:0")
	defer ln.Close()

	go server.Serve(ln)

	addr := ln.Addr().String()

	b.ReportAllocs()
	b.SetBytes(int64(len(payload)))
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			b.Fatal(err)
		}
		defer conn.Close()

		for pb.Next() {
			_ = conn.SetWriteDeadline(time.Now().Add(2 * time.Second))
			if _, err := conn.Write(payload); err != nil {
				if ne, ok := err.(net.Error); ok && ne.Timeout() {
					return
				}
				return
			}
		}
	})
}
