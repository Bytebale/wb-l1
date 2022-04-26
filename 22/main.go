package main

import (
	"fmt"
	"math/big"
	"os"
)

/*
Если передаваемая строка в setString() начинается с «0x», будет использоваться основание 16
(шестнадцатеричное). Если строка начинается с «0», будет использоваться восьмеричная система
счисления. В противном случае будет использоваться основание 10 (десятичное). Также можно указать
базу вручную (до 62). Возвращаемое значение имеет тип *big.Int. Если Go не удается создать big.Int,
для ok устанавливается значение false.
*/
func main() {
	result := new(big.Int)
	var a, ok = new(big.Int).SetString(os.Args[1], 0)
	var b, oke = new(big.Int).SetString(os.Args[2], 0)
	if !oke || !ok {
		fmt.Println("Error args")
	}

	result.Add(a, b)
	fmt.Println("Сложение: ", result)
	result.Sub(a, b)
	fmt.Println("Вычитание: ", result)
	result.Div(a, b)
	fmt.Println("Деление: ", result)
	result.Mul(a, b)
	fmt.Println("Умножение: ", result)
}
