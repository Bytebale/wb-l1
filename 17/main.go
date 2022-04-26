package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func binarySearch(arr []int, item int) int {
	low := 0 // определяем границы поиска
	high := len(arr) - 1

	for low <= high { // пока границы не сократились до одного
		mid := (low + high) // проверяем середину
		guess := arr[mid]
		if guess == item { // если нашли возвращаем индекс
			return mid
		}
		if guess > item { // если больше искомого числа
			high = mid - 1
		} else { // если меньше
			low = mid + 1
		}
	}
	return -1 // не нашли
}

func main() {
	// создаем рандомный набор чисел
	arr := make([]int, 5)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(7)
	}
	// сортируем его (обязательное условие для бинарного поиска)
	sort.Ints(arr)

	// ищем индекс в массиве, под которым храниться fn
	fn := 2
	indx := binarySearch(arr, fn)
	if indx == -1 {
		fmt.Println("Number not found")
		return
	}
	// выводим индекс искомого числа
	fmt.Printf("number %d at index %d\n", arr[indx], indx)
}
