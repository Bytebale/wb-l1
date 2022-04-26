package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func work(ch chan interface{}, cancel chan struct{}, n int) {
	for data := range ch {
		select {
		case <-ch:
			fmt.Printf("Worker No: %d made %v\n", n, data)
			time.Sleep(200 * time.Millisecond)
		case <-cancel:
			return
		}
	}
}

func main() {
	task := [...]interface{}{1, 'a', "bb", "ccc"}
	mainCh := make(chan interface{})
	cancel := make(chan struct{})
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	rand.Seed(time.Now().UnixNano())
	fmt.Println("Select number of workers: ")
	var n int
	fmt.Scan(&n)

	for n != 0 {
		go work(mainCh, cancel, n)
		n--
	}

	for {
		select {
		default:
			mainCh <- task[rand.Intn(len(task))]
		case <-sigChan:
			close(cancel)
			fmt.Println("finish work")
			return
		}
	}
}
