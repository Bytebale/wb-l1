package main

import (
	"fmt"
	"strings"
)

func main() {
	strf := "abccd"
	res := true
	m := make(map[string]struct{})

	str := strings.ToLower(strf)

	for _, c := range str {
		m[string(c)] = struct{}{}
	}

	if len(m) < len(strf) {
		res = false
	}

	fmt.Println(str+":", res)
}
