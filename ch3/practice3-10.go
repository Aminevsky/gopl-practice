package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("1234567890"))
	fmt.Println(comma("1230"))
	fmt.Println(comma("123"))
}


func comma(s string) string {
	n := len(s)

	if n <= 3 {
		return s
	}

	h := n % 3

	var buf bytes.Buffer
	if h > 0 {
		buf.WriteString(s[:h])
	}

	for i := h; i < n; {
		buf.WriteByte(',')
		buf.WriteString(s[i:i+3])
		i += 3
	}

	return buf.String()
}
