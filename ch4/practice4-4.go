package main

import "fmt"

func rotate(s []int, n int) []int {
	out := make([]int, 0)
	out = append(out, s[n:]...)
	out = append(out, s[0:n]...)
	return out
}

func main() {
	s := []int{0,1,2,3,4,5}
	res := rotate(s, 2)
	fmt.Println(res)
}
