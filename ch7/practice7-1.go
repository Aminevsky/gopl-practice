package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type WordCounter int
func (w *WordCounter) Write(p []byte) (int, error) {
	*w++
	return 1, nil
}

type LineCounter int
func (c *LineCounter) Write(p []byte) (int, error) {

}

func main() {
	file := os.Args[1]
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("Cannot open the file: %s", file)
	}
	defer f.Close()

	var w WordCounter

	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		line := input.Text()

		w.Write([]byte(line))
	}
	fmt.Println(w)
}
