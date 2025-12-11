package main

import (
	"fmt"
	"strings"
)
	
// Уникальные символы в строке

// Разработать программу, которая проверяет,
// что все символы в строке встречаются один раз
// (т.е. строка состоит из уникальных символов).

// Вывод:
// true, если все символы уникальны, false, если есть повторения.
// Проверка должна быть регистронезависимой,
// т.е. символы в разных регистрах считать одинаковыми.

// Например:
// "abcd" -> true,
// "abCdefAaf" -> false (повторяются a/A),
// "aabcd" -> false.

// Подумайте, какой структурой данных удобно
// воспользоваться для проверки условия.

func main() {
	fmt.Println(unique("aA"))
	// fmt.Println(byte('a')+32 - byte('A'))
}

// неэффективное решение, сложноность O(n)
func unique(str string) bool {
	seen := make(map[rune]bool)
	for _, char := range strings.ToLower(str) {
		if seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}

// неэффективное решение, сложноность O(n^2)
// func uniqueOld(str string) bool {
// 	counter := 0
// 	step := 0
// 	for counter < len(str) && step < len(str) {
// 		// fmt.Println(step, counter)
// 		if counter != step {
// 			if str[step] == str[counter] ||
// 				(str[step]+32 - str[counter]) == 0 ||
// 				(str[step]+32 - str[counter]) == 64 {
// 				return false
// 			}
// 		}
// 		if counter+1 == len(str) && step < len(str) {
// 			counter = 0
// 			step++
// 			continue
// 		}
// 		counter++
// 	}
// 	return true
// }