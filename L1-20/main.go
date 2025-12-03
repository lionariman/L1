package main

import (
	"fmt"
)

// Разворот слов в предложении

// Разработать программу, которая переворачивает порядок слов в строке.

// Пример: входная строка:
// «snow dog sun», выход: «sun dog snow».

// Считайте, что слова разделяются одиночным пробелом.
// Постарайтесь не использовать дополнительные срезы, а выполнять операцию «на месте».

func main() {
	res := reverseWords("sun dog snow")
	fmt.Printf("%v", res)
}

func reverseWords(words string) string {
	arr := []rune(words)
	reverse(arr, 0, len(words)-1)
	start := 0
	for i := 0; i <= len(arr); i++ {
		if i == len(arr) || arr[i] == ' ' {
			reverse(arr, start, i-1)
			start = i + 1
		}
	}
	return string(arr)
}

func reverse(runes []rune, left, right int) {
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
}

// First version

// func reverseWordsFirst(words string) string {
// 	arr := []rune{}
// 	left := len(words)-1
// 	right := len(words)-1
// 	for left > -1 {
// 		if words[left] == ' ' || left == 0 {
// 			left2 := left
// 			for left2 <= right {
// 				if rune(words[left2]) != 32 {
// 					arr = append(arr, rune(words[left2]))
// 				}
// 				left2++
// 			}
// 			if left > 0 {
// 				arr = append(arr, 32)
// 			}
// 			right = left
// 		}
// 		left--
// 	}
// 	return string(arr)
// }