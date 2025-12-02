package main

import (
	"fmt"
)

// Бинарный поиск

// Реализовать алгоритм бинарного поиска встроенными методами языка.
// Функция должна принимать отсортированный слайс и искомый элемент,
// возвращать индекс элемента или -1, если элемент не найден.

// Подсказка: можно реализовать рекурсивно или итеративно, используя цикл for.

func main() {
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 23, 42, 45, 60, 78, 90, 99}, 10))
}

func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr)-1
	for left <= right {
		mid := left + (right - left)/2
		if target < arr[mid] {
			right = mid-1
		} else if target > arr[mid] {
			left = mid+1
		} else {
			return mid
		}
	}
	return -1
}