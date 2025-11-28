package main

import "fmt"

// Собственное множество строк

// Имеется последовательность строк:
// ("cat", "cat", "dog", "cat", "tree").
// Создать для неё собственное множество.

// Ожидается: получить набор уникальных слов.
// Для примера, множество = {"cat", "dog", "tree"}.

func main() {
	fmt.Println("111")
	fmt.Println(leftUniqueStrings([]string{"cat", "cat", "dog", "cat", "tree"}))
}

func leftUniqueStrings(list []string) []string {
	mp := make(map[string]bool)
	st := []string{}
	for _, i := range list {
		mp[i] = true
	}
	
	for k := range mp {
		st = append(st, k)
	}
	
	return st
}