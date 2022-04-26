package main

import (
	"fmt"
	"time"
)

func wrt(ch chan int, val int) {
	ch <- val
}

func rd(ch chan int) {
	val := <-ch
	fmt.Println("Received data with value: ", val)
}

func main() {
	var t time.Duration // время на работу программы
	fmt.Print("Set time in seconds: ")
	fmt.Scan(&t)
	ch := make(chan int)
	timer := time.NewTimer(t * time.Second) // запускаем таймер с заданным временем
	i := 1

	for {
		select {
		case <-timer.C: // получаем сигнал по истечению времени
			fmt.Println("time is up")
			close(ch)
			return
		default:
			go wrt(ch, i)
			go rd(ch)
			time.Sleep(100 * time.Millisecond)
			i++
		}
	}
}
