package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	for _, arg := range os.Args[1:] {
		splitted := strings.Split(arg, "=")
		city := splitted[0]
		address := splitted[1]

		go fetch(city, address)
		time.Sleep(1 * time.Second)
	}
}

func fetch(city string, address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(city)
	mustCopy(os.Stdout, conn)

	conn.Close()
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
