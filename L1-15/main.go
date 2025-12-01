package main

import "fmt"

// Небольшой фрагмент кода — проблемы и решение

// Рассмотреть следующий код и ответить на вопросы:
// к каким негативным последствиям он может привести и как это исправить?
// Приведите корректный пример реализации.

// var justString string

// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }

// Вопрос: что происходит с переменной justString? 

// ОТВЕТ: при создании среза на 100 символов мы все также
// ссылаемся на ту же область памяти, которая была выделана
// на 1024 символов. чтобы исправить это надо создать новый срез
// который не будет ссылаться на прошлую область памяти.

func main() {
	v := createHugeString(1 << 10)
	justString := string([]rune(v[:100]))
	fmt.Println(len(justString))
}

func createHugeString(size int) string {
	hugeString := ""
	for i := 0; i < size; i++ {
		hugeString = hugeString + "0"
	}
	return hugeString
}