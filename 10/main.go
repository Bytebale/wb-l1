package main

import "fmt"

func main() {
	arr := []float64{-25.4, -27.0, 13.0, -9.0, 15.5, 24.5, -21.0, 32.5} // последовательность температур
	m := make(map[int][]float64)                                        // мапа где ключ нижний порог группы температур

	for _, temp := range arr {
		if temp < 0 {
			key := (int(temp) - 10) / 10 * 10 // изключение для группы от -10 до 0 градусов
			m[key] = append(m[key], temp)     // запись значений в мапу в соответсвии с ключом
		} else {
			key := (int(temp) / 10) * 10 // все остальные случаи
			m[key] = append(m[key], temp)
		}
	}
	for k, v := range m { // вывод в консоль в соотвеетсвии key-value
		fmt.Println(k, v)
	}
}
