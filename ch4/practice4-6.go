package main

import (
	"fmt"
	"unicode"
)

func duplicateSpace(s []byte) []byte {
	var spaces int
	for i:=0; i<len(s); {
		if unicode.IsSpace(rune(s[i])) {
			var j int
			for j=i+1; unicode.IsSpace(rune(s[j])); {
				j++
			}
			fmt.Printf("i=%d, j=%d\n", i, j)

			s[i] = byte(' ')
			copy(s[i+1:], s[j:])
			spaces += j - i
			i += j-i
		} else {
			i++
		}
	}

	return s[0:len(s)-spaces+1]
}

func main() {
	s := []byte{'A', '\t', '\t', 'B'}
	fmt.Println(s)
	fmt.Println(duplicateSpace(s))
	//fmt.Println(s)
}
