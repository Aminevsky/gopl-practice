package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func handleConn(conn net.Conn) {
	ch := make(chan string) // 送信用のクライアントメッセージ
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch // 到着を知らせる

	var isConnect = make(chan struct{})

	go func() {
		for {
			select {
			case <-time.After(10 * time.Second):
				fmt.Println("切断した")
				messages <- who + " has disconnected"
				conn.Close()
				return
			case <-isConnect:
				fmt.Println("受信した")
				continue
			}
		}
	}()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		isConnect <- struct{}{}
		messages <- who + ": " + input.Text()
	}

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
