package main

import (
	"fmt"
	"sync"
)

type pair struct {
	x int
	y int
}

func main() {
	pairs := []*pair{}
	pairChan := make(chan *pair)
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(val int, c chan *pair, w sync.WaitGroup) {
			p := handleNumber(val)
			fmt.Printf("%+v\n", p)
			pairs = append(pairs, p)
			w.Done()
		}(i, pairChan, wg)
	}
	wg.Wait()
	close(pairChan)
	for p := range pairChan {
		pairs = append(pairs, p)
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
