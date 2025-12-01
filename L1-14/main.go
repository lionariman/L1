package main

import (
	"fmt"
)

// Определение типа переменной в runtime

// Разработать программу, которая в runtime способна определить тип переменной,
// переданной в неё (на вход подаётся interface{}).
// Типы, которые нужно распознавать: int, string, bool, chan (канал).
// Подсказка: оператор типа switch v.(type) поможет в решении.

type Human struct {
	Name string
	Age int
}

func main() {
	// срез с разными значениями
	var values []interface{} = []interface{}{"312321", 123, 45.3, Human{"human", 23}}
	
	for _, val := range values {
		checkValType(val)
	}
}

func checkValType(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println(value, " is int")
	case string:
		fmt.Println(value, " is string")
	case Human:
		fmt.Println(value, " is Human")
	case float64:
		fmt.Println(value, " is float")
	default:
		fmt.Println("I do not know what type is this")
	}
}
