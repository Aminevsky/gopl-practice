package main

import (
	"strings"
	"testing"
	"unicode/utf8"
)

type testData struct {
	input  string
	result CharCount
}

func TestCountChar(t *testing.T) {
	res := CountChar(strings.NewReader("A"))
	if res.Counts['A'] != 1 {
		t.Errorf("CountChar(\"A\").Counts = %d", res.Counts['A'])
	}

	var tests = []testData{
		{
			input: "A",
			result: CharCount{
				Counts:  map[rune]int{'A': 1},
				Utflen:  [utf8.UTFMax + 1]int{1: 1},
				Invalid: 0,
				Err:     nil,
			},
		},
		{
			input: "AB",
			result: CharCount{
				Counts:  map[rune]int{'A': 1, 'B': 1},
				Utflen:  [utf8.UTFMax + 1]int{1: 2},
				Invalid: 0,
				Err:     nil,
			},
		},
		{
			input: "あ",
			result: CharCount{
				Counts:  map[rune]int{'あ': 1},
				Utflen:  [utf8.UTFMax + 1]int{3: 1},
				Invalid: 0,
				Err:     nil,
			},
		},
		{
			input: "あい",
			result: CharCount{
				Counts:  map[rune]int{'あ': 1, 'い': 1},
				Utflen:  [utf8.UTFMax + 1]int{3: 2},
				Invalid: 0,
				Err:     nil,
			},
		},
		{
			input: "Aあ",
			result: CharCount{
				Counts:  map[rune]int{'A': 1, 'あ': 1},
				Utflen:  [utf8.UTFMax + 1]int{1: 1, 3: 1},
				Invalid: 0,
				Err:     nil,
			},
		},
		{
			input: "",
			result: CharCount{
				Counts:  map[rune]int{},
				Utflen:  [utf8.UTFMax + 1]int{},
				Invalid: 0,
				Err:     nil,
			},
		},
		{
			input: "🍣",
			result: CharCount{
				Counts:  map[rune]int{'🍣': 1},
				Utflen:  [utf8.UTFMax + 1]int{4: 1},
				Invalid: 0,
				Err:     nil,
			},
		},
	}

	for _, d := range tests {
		res := CountChar(strings.NewReader(d.input))
		for _, ch := range d.input {
			if count := res.Counts[ch]; count != d.result.Counts[ch] {
				t.Errorf("Counts actual: %d, expected: %d", res.Counts[ch], d.result.Counts[ch])
			}
		}

		if res.Utflen != d.result.Utflen {
			t.Errorf("Utflen actual: %v, expected: %v", res.Utflen, d.result.Utflen)
		}

		if res.Invalid != d.result.Invalid {
			t.Errorf("Invalid actual: %d, expected: %d", res.Invalid, d.result.Invalid)
		}

		if res.Err != d.result.Err {
			t.Errorf("Error actual: %v, expected: %v", res.Err, d.result.Err)
		}
	}
}
