package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

type CharCount struct {
	Counts  map[rune]int
	Utflen  [utf8.UTFMax + 1]int
	Invalid int
	Err     error
}

func CountChar(r io.Reader) CharCount {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(r)

	for {
		r, n, err := in.ReadRune() // rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			return CharCount{Err: errors.New(fmt.Sprintf("charcount: %v\n", err))}
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}

	return CharCount{Counts: counts, Utflen: utflen, Invalid: invalid, Err: nil}
}

func main() {
	res := CountChar(os.Stdin)

	fmt.Printf("rune\tcount\n")
	for c, n := range res.Counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Print("\nlen\tcount\n")
	for i, n := range res.Utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	if res.Invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", res.Invalid)
	}
}
