package main

import "fmt"

func main() {
	mp := make(map[string]int)
	arr1 := []string{"tomat", "cucumber", "banan", "bread", "milk"}
	arr2 := []string{"bread", "milk"}
	arr3 := append(arr1, arr2...)
	var result []string

	for _, key := range arr3 {
		mp[key] += 1
		if mp[key] > 1 {
			result = append(result, key)
		}
	}
	fmt.Println("map: ", mp) // each map element with value 2 or more is intersection
	fmt.Println("res: ", result)
}
