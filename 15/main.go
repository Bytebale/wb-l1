package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"strings"
)

/*
Самый простой способ создать утечку памяти в Go — определить глобальную переменную,
массив, и добавить в него данные.

сверим выделяемую и удерживаемую память через "runtime/pprof"

	******************** исходный код  **********************************

      flat  flat%   sum%        cum   cum%
 1025.12kB 50.02% 50.02%  1025.12kB 50.02%  runtime.allocm
  512.38kB 25.00% 75.02%   512.38kB 25.00%  main.createHugeString
  512.04kB 24.98%   100%   512.04kB 24.98%  runtime.bgscavenge
         0     0%   100%   512.38kB 25.00%  main.main
         0     0%   100%   512.38kB 25.00%  main.someFunc
         0     0%   100%   512.56kB 25.01%  runtime.gopreempt_m
         0     0%   100%   512.56kB 25.01%  runtime.goschedImpl
         0     0%   100%   512.56kB 25.01%  runtime.handoffp
         0     0%   100%   512.38kB 25.00%  runtime.main
         0     0%   100%   512.56kB 25.01%  runtime.morestack

	********** без глобальной переменной *********************************

      flat  flat%   sum%        cum   cum%
 1537.69kB 50.01% 50.01%  1537.69kB 50.01%  runtime.allocm
 1536.92kB 49.99%   100%  1536.92kB 49.99%  main.createHugeString
         0     0%   100%  1536.92kB 49.99%  main.main
         0     0%   100%  1536.92kB 49.99%  main.someFunc
         0     0%   100%   512.56kB 16.67%  runtime.gopreempt_m
         0     0%   100%   512.56kB 16.67%  runtime.goschedImpl
         0     0%   100%   512.56kB 16.67%  runtime.handoffp
         0     0%   100%  1536.92kB 49.99%  runtime.main
         0     0%   100%   512.56kB 16.67%  runtime.morestack
         0     0%   100%  1025.12kB 33.34%  runtime.mstart

	********** без глобальной переменной с клонированием строки **********

      flat  flat%   sum%        cum   cum%
  512.56kB 50.02% 50.02%   512.56kB 50.02%  runtime.allocm
  512.20kB 49.98%   100%   512.20kB 49.98%  runtime.malg
         0     0%   100%   512.56kB 50.02%  runtime.mstart
         0     0%   100%   512.56kB 50.02%  runtime.mstart0
         0     0%   100%   512.56kB 50.02%  runtime.mstart1
         0     0%   100%   512.56kB 50.02%  runtime.newm
         0     0%   100%   512.20kB 49.98%  runtime.newproc.func1
         0     0%   100%   512.20kB 49.98%  runtime.newproc1
         0     0%   100%   512.56kB 50.02%  runtime.resetspinning
         0     0%   100%   512.56kB 50.02%  runtime.schedule

------ flat означает, что память выделена функцией и удерживается ей;
------ cum означает, что память выделена функцией или функцией, которая была вызвана стеком
*/

func createHugeString(size int) (result string) {
	for i := 0; i < size; i++ {
		result += "a"
	}
	fmt.Println("str: ", result)
	return result
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString := strings.Clone(v[:100])
	fmt.Println("v:   ", justString)
}

// var justString string

// func someFunc() {
// 	v := createHugeString(1 << 10)
// 	justString = v[:100]
// }

func main() {
	someFunc()
	heaps, err := os.Create("heap")
	if err != nil {
		fmt.Println(err)
	}
	alloc, err := os.Create("alloc")
	if err != nil {
		fmt.Println(err)
	}
	pprof.Lookup("heap").WriteTo(heaps, 0)
	pprof.Lookup("allocs").WriteTo(alloc, 0)
	alloc.Close()
	heaps.Close()
}
