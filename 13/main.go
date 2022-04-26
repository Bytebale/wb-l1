package main

import "fmt"

func main() {
	a := 21
	b := 42

	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)
}
