package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("golang.html")
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot open html file: %s", err)
	}

	doc, err := html.Parse(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot parse: %s", err)
	}

	displayTextNode(doc)
}

func displayTextNode(n *html.Node) {
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return
	}

	if n.Type == html.TextNode {
		trim := strings.TrimSpace(n.Data)
		if trim != "" {
			fmt.Println(trim)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		displayTextNode(c)
	}
}
