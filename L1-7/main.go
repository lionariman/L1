package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// Реализовать безопасную для конкуренции запись данных в структуру map.
// Подсказка: необходимость использования синхронизации (например, sync.Mutex или встроенная concurrent-map).
// Проверьте работу кода на гонки (util go run -race).

func main() {
	// concurrentWriteWithMutex()
	concurrentWriteWithSyncMap()
}

func concurrentWriteWithMutex() {
	mt := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	orders := make(map[int]string, 10)

	wg.Add(10)

	for i := range 10 {
		go func(i int) {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			mt.Lock()
			orders[i] = "order-" + strconv.Itoa(i)
			fmt.Println("G", i, " -> ", orders[i])
			mt.Unlock()
		}(i)
	}

	wg.Wait()
}

func concurrentWriteWithSyncMap() {
	var orders sync.Map
	wg := &sync.WaitGroup{}

	wg.Add(10)

	for i := range 10 {
		go func(i int) {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			orders.Store(i, "order-"+strconv.Itoa(i))
			fmt.Println("G->", i)
		}(i)
	}

	wg.Wait()
}
