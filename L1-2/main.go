package main

import (
	"fmt"
	"sync"
)

func main() {

	mt := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	numbers := [5]int{2, 4, 6, 8, 10}

	wg.Add(len(numbers))

	for i := 0; i < len(numbers); i++ {
		go func() {
			mt.Lock()
			wg.Done()
			squareOfNum := numbers[i] * numbers[i]
			fmt.Println("Goroutine", i, "> квадрат числа", numbers[i], "=", squareOfNum)
			numbers[i] = squareOfNum
			mt.Unlock()
		}()
	}

	wg.Wait()

	// fmt.Println(numbers)
}
