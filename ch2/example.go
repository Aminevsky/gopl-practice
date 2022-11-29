package main

import (
	"fmt"
	"gopl/ch2/popcount"
	"time"
)

func main() {
	popcount1()
	popcount2()
	popcount3()
}

func popcount1() {
	start := time.Now()
	res := popcount.PopCount(2632875912)
	fmt.Printf("%ds elapsed\n", time.Since(start).Milliseconds())
	fmt.Printf("res: %d\n", res)
}

func popcount2() {
	start := time.Now()
	res := popcount.PopCount2(2632875912)
	fmt.Printf("%ds elapsed\n", time.Since(start).Milliseconds())
	fmt.Printf("res: %d\n", res)
}

func popcount3() {
	start := time.Now()
	res := popcount.PopCount3(2632875912)
	fmt.Printf("%ds elapsed\n", time.Since(start).Milliseconds())
	fmt.Printf("res: %d\n", res)
}
