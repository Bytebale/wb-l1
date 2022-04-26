package main

import (
	"fmt"
)

func removeElem(s []int, i int) []int {
	sl := append(s[:i], s[i+1:]...) // создаем новый срез без i-того элемента
	return sl
}

func removeFast(s []int, i int) []int {
	s[i] = s[len(s)-1]  // последний элемент на место i-того
	return s[:len(s)-1] // возвращаем без последнего элемента
}

func main() {
	slice := []int{2, 5, 7, 9, 11, 13}
	slice2 := []int{2, 5, 7, 9, 11, 13}
	i := 3
	fmt.Println("initial slice:    ", slice)
	fmt.Println("after removeElem: ", removeElem(slice, i))  // удаление с сохранением порядка
	fmt.Println("after removefast: ", removeFast(slice2, i)) // без сохранения порядка
}
