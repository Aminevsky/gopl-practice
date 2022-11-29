package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	var wg sync.WaitGroup

	tcpConn, _ := c.(*net.TCPConn)
	input := bufio.NewScanner(tcpConn)
	for input.Scan() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			echo(tcpConn, input.Text(), 1*time.Second)
		}()
	}

	wg.Wait()
	tcpConn.CloseWrite()
}
