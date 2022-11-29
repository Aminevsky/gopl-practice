package main

import "fmt"

func max(a int, b ...int) int {
	if len(b) == 0 {
		return a
	}

	n := []int{a}
	n = append(n, b...)

	tmp := a
	for _, v := range n {
		if v > tmp {
			tmp = v
		}
	}

	return tmp
}

func min(a int, b ...int) int {
	if len(b) == 0 {
		return a
	}

	n := []int{a}
	n = append(n, b...)

	tmp := a
	for _, v := range n {
		if v < tmp {
			tmp = v
		}
	}

	return tmp
}

func main() {
	fmt.Println(max(1))
	fmt.Println(max(1, 2))
	fmt.Println(max(1, 2, 3))

	fmt.Println(min(10))
	fmt.Println(min(9, 10))
	fmt.Println(min(8, 9, 10))
}
