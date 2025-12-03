package main

import (
	"fmt"
	// "math/big"
)
	
// Большие числа и операции

// Разработать программу, которая перемножает, делит, складывает, вычитает
// две числовых переменных a, b, значения которых > 2^20 (больше 1 миллион).

// Комментарий: в Go тип int справится с такими числами,
// но обратите внимание на возможное переполнение для ещё больших значений.
// Для очень больших чисел можно использовать math/big.

func main() {
	fmt.Println(calculator(1, 2, '+'))
}

func calculator(a, b int, action rune) int {
	switch action {
		case '+':
			return a + b
		case '-':
			return a - b
		case '*':
			return a * b
		case '/':
			return a / b
		default:
			fmt.Println("Not found action")
	}
	return 0
}