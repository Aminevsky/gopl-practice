package main

import "fmt"

func Join(sep string, elems ...string) string {
	if len(elems) == 0 {
		return ""
	}
	if len(elems) == 1 {
		return elems[0]
	}

	str := elems[0]
	for _, elem := range elems[1:] {
		str += sep
		str += elem
	}

	return str
}

func main() {
	fmt.Println(Join(", "))
	fmt.Println(Join(", ", "A"))
	fmt.Println(Join(", ", "A", "B"))
	fmt.Println(Join(", ", "A", "B", "C"))
}
