package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	split_str := strings.Split(str, " ")
	result := ""
	for _, word := range split_str {
		result = word + " " + result
	}
	fmt.Print(str + "-" + result + "\n")
}
