package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	// "time"
)

// Конкурентный счетчик

// Реализовать структуру-счётчик, которая будет инкрементироваться
// в конкурентной среде (т.е. из нескольких горутин).
// По завершению программы структура должна выводить итоговое значение счётчика.

// Подсказка: вам понадобится механизм синхронизации,
// например, sync.Mutex или sync/Atomic для безопасного инкремента.

type Counter struct {
	Count atomic.Int64
}

func (cn *Counter) Increment() {
	cn.Count.Add(1)
}

func main() {
	counter := Counter{}
	
	// mt := sync.Mutex{}
	wg := sync.WaitGroup{}
	
	gnum := 3
	
	wg.Add(gnum)
	
	for i := range gnum {
		go func(j int, cnt *Counter) {
			defer wg.Done()
			// mt.Lock()
			cnt.Increment()
			fmt.Println("Goroutine", j, " made increment >", cnt.Count.Load())
			// time.Sleep(1 * time.Second)
			// mt.Unlock()
		}(i, &counter)
	}
	
	wg.Wait()
	
	fmt.Println("Result of incrementation:", counter.Count.Load())
}