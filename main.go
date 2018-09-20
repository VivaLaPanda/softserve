package main

import (
	"bufio"
	"io"
	"net"
	"net/http"
	"os"
)

func handle(w http.ResponseWriter, req *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		panic("expected http.ResponseWriter to be an http.Flusher")
	}
	w.Header().Set("Transfer-Encoding", "chunked")
	var err error
	reader := bufio.NewReader(os.Stdin)
	for i := 1; err == nil; i++ {
		_, err = io.CopyN(w, reader, 32)
		flusher.Flush() // Trigger "chunked" encoding and send a chunk...
	}
}

func main() {
	/* Net listener */
	n := "tcp"
	addr := "127.0.0.1:9099"
	l, err := net.Listen(n, addr)
	if err != nil {
		panic("Failed to start server")
	}

	/* HTTP server */
	server := http.Server{
		Handler: http.HandlerFunc(handle),
	}
	server.Serve(l)
}
