package main

import "fmt"
	
// Удаление элемента слайса

// Удалить i-ый элемент из слайса.
// Продемонстрируйте корректное удаление без утечки памяти.

// Подсказка:
// можно сдвинуть хвост слайса на место удаляемого элемента
// (copy(slice[i:], slice[i+1:])) и уменьшить длину слайса на 1.

func main() {
	arr := []int{3, 4, 9, 123, 55, 20, 754, 8, 2, 6}
	fmt.Println(arr)
	deleteByIndex2(&arr, 5)
	fmt.Println(arr)
}

// мой вариант
func deleteByIndex(arr *[]int, ind int) {
	ln := len(*arr)
	for i := 0; i < ln; i++ {
		if i >= ind && i < ln-1 {
			(*arr)[i] = (*arr)[i+1]
		}
	}
	(*arr)[ln-1] = 0
	*arr = (*arr)[:ln-1]
}

// Более правильный вариант
func deleteByIndex2(arr *[]int, ind int) {
	copy((*arr)[ind:], (*arr)[ind+1:])
	(*arr)[len(*arr)-1] = 0
	*arr = (*arr)[:len(*arr)-1]
}