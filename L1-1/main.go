package main

import "fmt"

// Тут показываем как одну структуру можно засунуть в другую
type Human struct {
	Name  string
	Age   int
	Mail  string
	Phone int
}

// А тут уже Action "наследует" все поля от Human
// Просто пишем Human без имени поля - и все его поля становятся доступны
type Action struct {
	Human
	field string
}

func main() {
	// Создаем Action, внутри которого есть Human + свое поле
	// Теперь у action1 есть все поля от Human (Name, Age, Mail, Phone) + field
	action1 := Action{Human{"Name1", 23, "name1@gmail.com", 89674039172}, "field1"}

	fmt.Println("action1:", action1)
	// Можем обращаться к полям Human напрямую: action1.Name, action1.Age и т.д.
}
