package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		panic("error")
	}

	go func() {
		io.Copy(os.Stdout, tcpConn)
		log.Println("done")
		done <- struct{}{}
	}()

	mustCopy(tcpConn, os.Stdin)

	<-done
	tcpConn.CloseRead()
	tcpConn.CloseWrite()
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
