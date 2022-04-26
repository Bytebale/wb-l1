package main

import (
	"fmt"
	"reflect"
)

/*
reflect.TypeOf returns the reflection Type that represents the dynamic type of i. If i is a nil
interface value, TypeOf returns nil.
*/

func main() {
	var t interface{}
	t = 1
	fmt.Println(reflect.TypeOf(t))
	t = "qwerty"
	fmt.Println(reflect.TypeOf(t))
	t = true
	fmt.Println(reflect.TypeOf(t))
	t = make(chan struct{})
	fmt.Println(reflect.TypeOf(t))
}
