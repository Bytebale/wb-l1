package main

import (
	"fmt"
	"sync"
)

// счетчик с мьютексом
type counter struct {
	sync.Mutex
	counter int
}

// конструктор
func newCount() *counter {
	return &counter{}
}

// инкремент в конкурентной среде
func (c *counter) increment() {
	c.Lock()
	c.counter++
	c.Unlock()
}

func main() {
	counter := newCount()
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			counter.increment()
			fmt.Println("inc: ", counter.counter)
			wg.Done()
		}(wg)
	}
	wg.Wait()
	fmt.Println("total: ", counter.counter)
}
