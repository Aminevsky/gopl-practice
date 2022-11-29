package main

import (
	"fmt"
	"log"
	"os"
)

var cwd string

func init() {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
}

func main() {
	fmt.Println("started")
	fmt.Printf("%s\n", cwd)
}
