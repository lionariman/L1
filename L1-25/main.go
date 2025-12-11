package main

import (
	"fmt"
	"time"
	"sync"
)

// Своя функция Sleep

// Реализовать собственную функцию sleep(duration)
// аналогично встроенной функции time.Sleep,
// которая приостанавливает выполнение текущей горутины.

// Важно:
// в отличие от настоящей time.Sleep,
// ваша функция должна именно блокировать выполнение
// (например, через таймер или цикл),
// а не просто вызывать time.Sleep :) — это упражнение.

// Можно использовать канал + горутину,
// или цикл на проверку времени
// (не лучший способ, но для обучения).

func sleep(d time.Duration) {
	target := time.Now().UnixNano() + d.Nanoseconds()
	for {
		if time.Now().UnixNano() >= target {
			return
		}
	}
}

func main() {
	wg := sync.WaitGroup{}
	
	wg.Add(2)
	
	go func() {
		defer wg.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("default sleep")
	}()
	go func() {
		defer wg.Done()
		sleep(5 * time.Second)
		fmt.Println("my sleep")
	}()
	
	wg.Wait()
	
	fmt.Println("yes")
}