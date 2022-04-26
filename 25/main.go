package main

import (
	"fmt"
	"time"
)

func ft_sleep(t int) {
	<-time.After(time.Second * time.Duration(t))
}

func main() {
	fmt.Println(123)
	ft_sleep(2)
	fmt.Println(321)
}
