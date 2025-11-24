package main

import "fmt"

// Пересечение множеств

// Реализовать пересечение двух неупорядоченных множеств
// (например, двух слайсов) — т.е. вывести элементы,
// присутствующие и в первом, и во втором.

// Пример:
// A = {1,2,3}
// B = {2,3,4}
// Пересечение = {2,3}

func main() {
	a := []int{1, 2, 3, 7, 8, 9}
	b := []int{2, 3, 4, 9, 8, 7}
	x := intersection(a, b)
	fmt.Println(x)
}

func intersection(a []int, b []int) []int {
	setA := make(map[int]bool)
	for _, v := range a {
		setA[v] = true
	}

	result := []int{}
	seen := make(map[int]bool) // для избежания дубликатов

	for _, v := range b {
		if setA[v] && !seen[v] {
			result = append(result, v)
			seen[v] = true
		}
	}

	return result
}
