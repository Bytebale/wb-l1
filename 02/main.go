package main

import (
	"fmt"
	"sync"
)

func pow(i int, wg *sync.WaitGroup) {
	pow := i * i
	wg.Done()
	fmt.Printf("Square %d is %d\n", i, pow)
}

func main() {
	var num = []int{2, 4, 6, 8, 10}
	wg := &sync.WaitGroup{}
	for _, i := range num {
		wg.Add(1)
		go pow(i, wg)
	}
	wg.Wait()
}
