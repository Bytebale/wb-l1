package main

import (
	"fmt"
	"sync"
)

var mut = &sync.Mutex{}
var wg = &sync.WaitGroup{}
var m = make(map[int]int)

func wrMap(wg *sync.WaitGroup, data int, i int) {
	defer wg.Done()
	mut.Lock()
	m[i] = data
	mut.Unlock()
}

func main() {
	data := 6

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go wrMap(wg, data, i)
	}
	wg.Wait()

	for k, v := range m {
		fmt.Println("key:", k, " : ", "value", v)
	}

}
