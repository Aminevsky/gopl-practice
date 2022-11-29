package main

import "fmt"

func main() {
	var funcs []func()

	items := []int{1,2,3,4,5}
	for _, n := range items {
		funcs = append(funcs, func() {
			printNumber(n)
		})
	}

	for _, printFunc := range funcs {
		printFunc()
	}
}

func printNumber(n int) {
	fmt.Println(n)
}
