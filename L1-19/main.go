package main

import (
	"fmt"
)

// Разворот строки

// Разработать программу, которая переворачивает подаваемую на вход строку.
// Например: при вводе строки «главрыба» вывод должен быть «абырвалг».

// Учтите, что символы могут быть в Unicode (русские буквы, emoji и пр.),
// то есть просто iterating по байтам может не подойти — нужен срез рун ([]rune).

func main() {
	fmt.Println(
		stringReverse("главрыба"),"\n",
		stringReverse("Hello🌍"),"\n",
		stringReverse("1234567890qwertyuiop"))
}

func stringReverse(str string) string {
	runes := []rune(str)
	left := 0
	right := len(runes)-1
	for left < right {
		saving := runes[left]
		runes[left] = runes[right]
		runes[right] = saving
		left++
		right--
	}
	return string(runes)
}