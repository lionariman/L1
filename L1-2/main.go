package main

import (
	"fmt"
	"sync"
)

// Тут показываем как несколько горутин могут работать параллельно
// Но при этом не мешать друг другу благодаря блокировке через мьютекс

func main() {
	mt := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	numbers := [5]int{2, 4, 6, 8, 10}

	// Говорим WaitGroup что будем ждать 5 горутин
	wg.Add(len(numbers))

	// Запускаем 5 горутин (по одной для каждого числа)
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

	// Ждем пока все горутины закончат работу
	wg.Wait()

	// fmt.Println(numbers)
}
