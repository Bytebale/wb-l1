package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	quit := make(chan struct{})
	arr := []int{1, 2, 3, 5, 6, 7, 8, 9, 10, 175} //array with numbers

	go func(ch1 chan int) {
		for i := range arr {
			ch1 <- arr[i] // get numbers from array
		}
		close(ch1)
	}(ch1)

	go func(ch1, ch2 chan int) {
		for j := range ch1 {
			ch2 <- j * 2
		}
		close(ch2)
	}(ch1, ch2)

	go func(ch2 chan int, quit chan struct{}) {
		for val := range ch2 {
			fmt.Println(val)
		}
		quit <- struct{}{}
	}(ch2, quit)
	<-quit
}
