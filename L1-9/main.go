package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

// Разработать конвейер чисел.
// Даны два канала: в первый пишутся числа x из массива,
// во второй – результат операции x*2.
// После этого данные из второго канала должны выводиться в stdout.
// То есть, организуйте конвейер из двух этапов с горутинами:
// генерация чисел и их обработка.
// Убедитесь, что чтение из второго канала корректно завершается.

func main() {
	wg := sync.WaitGroup{}
	chanNumbers := make(chan int, 10)
	chanResults := make(chan int, 10)

	wg.Add(2)
	go genNumGoroutine(&wg, chanNumbers, 10)
	go mulNumGoroutine(&wg, chanNumbers, chanResults)
	wg.Wait()

	for num := range chanResults {
		fmt.Println(num)
	}
}

func genNumGoroutine(wg *sync.WaitGroup, readChan chan<- int, chanLen int) {
	defer wg.Done()
	for range chanLen {
		readChan <- rand.IntN(chanLen)
	}
	close(readChan)
}

func mulNumGoroutine(wg *sync.WaitGroup, readChan <-chan int, writeChan chan<- int) {
	defer wg.Done()
	for num := range readChan {
		writeChan <- (num * 2)
	}
	close(writeChan)
}
