package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] &= 0 << bit
}

func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')

	for i, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}

	buf.WriteByte('}')

	return buf.String()
}

func (s *IntSet) Len() int {
	elems := 0
	fmt.Println(s.words)
	for i, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if 64*i+j >= 0 { // これが無いとおかしくなる
					elems++
				}
			}
		}
	}
	return elems
}

func (s *IntSet) Copy() *IntSet {
	return &IntSet{words: s.words}
}

func (s *IntSet) AddAll(params ...int) {
	for _, v := range params {
		s.Add(v)
	}
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())

	//x.UnionWith(&y)
	//fmt.Println(x.String())

	//fmt.Println(x.Has(9), x.Has(123))

	fmt.Printf("x: %d, y: %d\n", x.Len(), y.Len())

	x.Remove(144)
	fmt.Println(x.String())

	x.Clear()
	fmt.Println(x.String())
	fmt.Printf("%d\n", x.Len())

	z := y.Copy()
	fmt.Println(z.String())
	z.Add(80)
	z.Add(90)
	z.Add(100)
	fmt.Println(z.String())
	fmt.Println(y.String())

	x.AddAll(10, 20, 30, 80, 90)
	fmt.Println(x.String())

	x.IntersectWith(z)
	fmt.Println(x.String())

	x.AddAll(10, 20, 30, 80, 90)
	fmt.Println(x.String())
	fmt.Println(z.String())
	x.DifferenceWith(z)
	fmt.Println(x.String())

	x.AddAll(10, 20, 30, 80, 90)
	fmt.Println(x.String())
	fmt.Println(z.String())
	x.SymmetricDifference(z)
	fmt.Println(x.String())
}
