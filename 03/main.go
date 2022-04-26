package main

import (
	"fmt"
	"sync"
)

func pow(i int, ch chan int, wg *sync.WaitGroup) {
	pow := i * i
	wg.Done()
	ch <- pow
}

func main() {
	var num = []int{2, 4, 6, 8, 10}
	var sum int
	ch := make(chan int, 1)
	wg := &sync.WaitGroup{}
	for _, i := range num {
		wg.Add(1)
		go pow(i, ch, wg)
	}
	wg.Wait()
	for i := 0; i < len(num); i++ {
		sum += <-ch
		fmt.Printf("Sum is n : %d\n", sum)
	}
	close(ch)
	fmt.Printf("Sum is : %d\n", sum)
}
