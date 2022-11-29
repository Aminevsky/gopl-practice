package main

import "fmt"

var x, y int

func main() {
	done := make(chan struct{})
	go func() {
		x = 1
		fmt.Print("y:", y, " ")
		done <- struct{}{}
	}()

	go func() {
		y = 1
		fmt.Print("x:", x, " ")
		done <- struct{}{}
	}()

	for i := 0; i < 2; i++ {
		<-done
	}
}
