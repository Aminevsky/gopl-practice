package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

type Counts map[string]int

func main() {
	f, err := os.Open("golang.html")
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot open html file: %s", err)
	}

	doc, err := html.Parse(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot parse: %s", err)
	}

	counts := make(map[string]int)
	countElement(counts, doc)

	for k, v := range counts {
		fmt.Printf("%s\t%d\n", k, v)
	}
}

func countElement(counts Counts, n *html.Node) {
	if n.Type == html.ElementNode {
		counts[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		countElement(counts, c)
	}
}
