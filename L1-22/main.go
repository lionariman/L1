package main

import (
	"fmt"
	"math/big"
)
	
// Большие числа и операции

// Разработать программу, которая перемножает, делит, складывает, вычитает
// две числовых переменных a, b, значения которых > 2^20 (больше 1 миллион).

// Комментарий: в Go тип int справится с такими числами,
// но обратите внимание на возможное переполнение для ещё больших значений.
// Для очень больших чисел можно использовать math/big.

func main() {
	fmt.Println(calculator(1_234_234_123_234, 2_123_345_123_897, '*'))
}

func calculator(num1, num2 int64, action rune) *big.Int {
	a := big.NewInt(num1)
	b := big.NewInt(num2)
	switch action {
		case '+':
			return a.Add(a, b)
		case '-':
			return a.Sub(a, b)
		case '*':
			return a.Mul(a, b)
		case '/':
			if num2 == 0 {
				panic("division by zero")
			}
			return a.Div(a, b)
		default:
			panic("unknown operation")
	}
}