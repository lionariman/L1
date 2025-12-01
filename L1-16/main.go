package main

import "fmt"

// Быстрая сортировка (quicksort)

// Реализовать алгоритм быстрой сортировки массива встроенными средствами языка.
// Можно использовать рекурсию.

// Подсказка: напишите функцию quickSort([]int) []int
// которая сортирует срез целых чисел.
// Для выбора опорного элемента можно взять середину или первый элемент.
	
func main() {
	unsortedNumbers := []int{6, 3, 8, 1, 4, 7, 2}
	sortedNumbers := quickSort(unsortedNumbers)
	
	fmt.Println(unsortedNumbers, "\n", sortedNumbers)
}

func quickSort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}
	pivot := len(numbers)/2
	left := []int{}
	right := []int{numbers[pivot]}
	for i, val := range numbers {
		if i == pivot {
			continue
		}
		if val < numbers[pivot] {
			left = append(left, val)
		} else {
			right = append(right, val)
		}
	}
	
	left = quickSort(left)
	right = quickSort(right)
	
	return append(left, right...)
}