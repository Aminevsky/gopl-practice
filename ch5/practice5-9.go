package main

import (
	"fmt"
	"strings"
)

func main() {
	before := "NewYork Tokyo Beijing $foo Paris"
	after := expand(before, converter)
	fmt.Println(after)
}

func converter(s string) string {
	return "Taipei"
}

func expand(s string, f func(string) string) string {
	words := strings.Split(s, " ")

	for i, word := range words {
		if word[0] == '$' {
			param := word[1:]
			words[i] = f(param)
		}
	}

	return strings.Join(words, " ")
}
