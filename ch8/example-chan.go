package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("処理結果: %s\n", mirroredQuery())
}

func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() { responses <- request("a") }()
	go func() { responses <- request("b") }()
	go func() { responses <- request("c") }()
	return <-responses
}

func request(hostname string) (responses string) {
	if hostname == "a" {
		time.Sleep(30 * time.Second)
	} else if hostname == "b" {
		time.Sleep(10 * time.Second)
	} else {
		time.Sleep(60 * time.Second)
	}

	return fmt.Sprintf("%s: Success", hostname)
}
