package main

import "fmt"

// Структура Human
type Human struct {
	age uint
}

// При передачи структуры Human в струкстуру Action, Action наследует данные и методы от Human
type Action struct {
	Human
	name string
}

func (h Human) getData() uint {
	return h.age
}

func main() {

	var name string
	var age uint
	fmt.Println("What is your name?")
	fmt.Scanf("%s", &name)
	fmt.Println("How old are you?")
	fmt.Scanf("%d", &age)

	h := Human{age: age}
	a := Action{Human: h, name: name}
	fmt.Printf("My name is %s I'm %d years old", a.name, a.getData())
}
