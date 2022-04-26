package main

import "fmt"

/*
Go не имеет явных конструкторов. Идиоматический способ установки новых
структур данных - это использование правильных нулевых значений в сочетании
с фабричными функциями (factory functions).
*/

type Point struct {
	x, y int
}

// Фабричная функция
func NewPoint(x int, y int) Point {
	return Point{x: x, y: y}
}

func (p *Point) getDistance() int {
	// исключили отрицательное значение в расстоянии
	if p.x < p.y {
		return p.y - p.x
	}
	// получили расстояние
	return p.x - p.y
}

func main() {
	var x, y int
	fmt.Println("set x, y: ")
	fmt.Scan(&x)
	fmt.Scan(&y)
	point := NewPoint(x, y)
	fmt.Println("distance: ", point.getDistance())
}
