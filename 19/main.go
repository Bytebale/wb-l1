package main

import "fmt"

/*
To convert a string into array of characters (which is slice of runes) in Go language,
pass the string as argument to []rune(). This will create a slice of runes where each
character is stored as an element in the resulting slice
*/

func main() {
	str := "главрыба"
	var revStr string
	chars := []rune(str)

	for i := len(chars) - 1; i >= 0; i-- {
		revStr += string(chars[i])
	}
	fmt.Println(str, "-", revStr)
}
