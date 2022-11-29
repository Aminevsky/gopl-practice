package main

import "fmt"

func main() {
	//var stack []int
	//stack = append(stack, 1, 2, 3)
	//fmt.Printf("%d\n", stack)
	//
	//top := stack[len(stack)-1]
	//fmt.Printf("%d\n", top)
	//
	//stack = stack[:len(stack)-1]
	//fmt.Printf("%d\n", stack)

	s := []int{5,6,7,8,9}
	fmt.Println(remove(s, 2))
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
