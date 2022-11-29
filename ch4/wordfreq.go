package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: wordfreq [file]")
	}

	file := os.Args[1]
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("Cannot open the file: %v\n", err)
	}
	defer f.Close()

	wordCounts := make(map[string]int)

	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		word = strings.Trim(word, ".,")
		wordCounts[word]++
	}

	for k, v := range wordCounts {
		fmt.Printf("%s\t%d\n", k, v)
	}
}
