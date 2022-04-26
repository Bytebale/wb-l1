package main

import (
	"fmt"
	"strconv"
)

func main() {
	var data int64
	var pos int64 // position of bit
	var v int64   // 1 || 0
	data = 1      // initial data

	fmt.Print("Enter position: ")
	fmt.Scan(&pos)
	fmt.Print("Enter 1 or 0: ")
	fmt.Scan(&v)

	if v == 0 {
		fmt.Println("Initial:", strconv.FormatInt(data, 2))                              // initial data in binary
		fmt.Println("Result: ", strconv.FormatInt(data&(9223372036854775807-1<<pos), 2)) // set 0 on position in binary
	}
	if v == 1 {
		fmt.Println("Initial:", strconv.FormatInt(data, 2))        // initial data in binary
		fmt.Println("Result: ", strconv.FormatInt(data|1<<pos, 2)) // set 1 on position in binary
	}
}
