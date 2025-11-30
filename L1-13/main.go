package main

import "fmt"

// Обмен значениями без третьей переменной

// Поменять местами два числа без использования временной переменной.
// Подсказка: примените сложение/вычитание или XOR-обмен.

func main() {
	vr1 := 123
	vr2 := 456

	fmt.Println("SWAP BY ADD/SUBSTRUCT")
	fmt.Println("BEFORE:", vr1, vr2)
	swapByAddSubstruct(&vr1, &vr2)
	fmt.Println(" AFTER:", vr1, vr2)
	
	fmt.Println("\nSWAP BY XOR")
	fmt.Println("BEFORE:", vr1, vr2)
	swapByXOR(&vr1, &vr2)
	fmt.Println(" AFTER:", vr1, vr2)
	
}

func swapByAddSubstruct(num1, num2 *int) {
	*num1 = *num1 + *num2
	*num2 = *num1 - *num2
	*num1 = *num1 - *num2
}

func swapByXOR(num1, num2 *int) {
	*num1 = *num1 ^ *num2
	*num2 = *num1 ^ *num2
	*num1 = *num1 ^ *num2
}
