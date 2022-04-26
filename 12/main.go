package main

import "fmt"

func main() {
	str := []string{"cat", "cat", "dog", "cat", "tree"}
	mp := make(map[string]int)
	var res []string

	for _, key := range str {
		mp[key] += 1
		if mp[key] == 1 {
			res = append(res, key)
		}
	}
	fmt.Println("map: ", mp)
	fmt.Println("res: ", res)
}
