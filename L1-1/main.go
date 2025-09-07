package main

import "fmt"

type Human struct {
	Name  string
	Age   int
	Mail  string
	Phone int
}

type Action struct {
	Human
	field string
}

func main() {
	action1 := Action{Human{"Name1", 23, "name1@gmail.com", 89674039172}, "field1"}
	fmt.Println("action1:", action1)
}
