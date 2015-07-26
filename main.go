package main

import (
	"fmt"
)

type pair struct {
	x int
	y int
}

func main() {
	pairs := []*pair{}
	for i := 0; i < 10; i++ {
		p := handleNumber(i)
		fmt.Printf("%+v\n", p)
		pairs = append(pairs, p)
		fmt.Println("looping")
	}
	fmt.Println("Done")
}

func handleNumber(i int) *pair {
	val := i
	if i%2 == 0 {
		val = f(i)
	}
	return &pair{
		x: i,
		y: val,
	}
}

func f(x int) int {
	return x*x + x
}
