package main

import "fmt"

type IntList struct {
	Value int
	Tail *IntList
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}

	return list.Value + list.Tail.Sum()
}

func main() {
	list := IntList{Value: 1, Tail: &IntList{2, nil}}
	result := list.Sum()
	fmt.Println(result)
}
