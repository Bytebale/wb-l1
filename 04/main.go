package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func work(ch chan interface{}, wg *sync.WaitGroup, n int) {
	for data := range ch {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("Worker No: %d made %v\n", n, data)
	}
	wg.Done()
}

func main() {
	task := [...]interface{}{1, 'a', "bb", "ccc"}
	mainCh := make(chan interface{})
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Select number of workers: ")
	var n int
	fmt.Scan(&n)
	wg := &sync.WaitGroup{}

	for n != 0 {
		wg.Add(1)
		go work(mainCh, wg, n)
		n--
	}

	for {
		select {
		default:
			mainCh <- task[rand.Intn(len(task))]
		case <-sigChan:
			close(mainCh)
			return
		}
	}
}
