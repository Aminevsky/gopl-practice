package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	typeFlag := flag.String("t", "sha256", "specify digest type")
	flag.Parse()

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		param := input.Text()

		if *typeFlag == "sha256" {
			digest := sha256.Sum256([]byte(param))
			fmt.Printf("%x\n", digest)
		}

		if *typeFlag == "sha384" {
			digest := sha512.Sum384([]byte(param))
			fmt.Printf("%x\n", digest)
		}

		if *typeFlag == "sha512" {
			digest := sha512.Sum512([]byte(param))
			fmt.Printf("%x\n", digest)
		}
	}
}
