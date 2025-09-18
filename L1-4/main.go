package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

// В этой задаче учимся Graceful shutdown - красиво завершать работу горутин
// Когда пользователь жмет Ctrl+C, мы не просто убиваем программу
// А вежливо просим всех воркеров завершить свою работу и дождаемся их
// Тут используем канал для оповещения (можно было еще context использовать)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Error: wrong number of args")
		return
	}
	numOfWorkers, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Error: converting error")
		return
	}

	fmt.Println("Number of Workers:", numOfWorkers)

	// c - это канал для получения сигналов от операционки (Ctrl+C)
	c := make(chan os.Signal, 1)
	// chain - это канал для оповещения воркеров о том что пора завершаться
	chain := make(chan struct{})
	// Слушаем ОС - когда пользователь нажмет Ctrl+C, она положит сигнал в канал c
	signal.Notify(c, os.Interrupt)

	wg := sync.WaitGroup{}
	wg.Add(numOfWorkers)

	for i := range numOfWorkers {
		go func(i int) {
			defer wg.Done()
			for {
				// select - чтобы горутина могла выполнять свою работу и также слушать канал
				// без него не получается - горутина просто зависает слушая канал
				select {
				case <-chain: // Если кто-то положил что-то в chain (сигнал о завершении)
					fmt.Println("g[", i, "] is closed")
					return // выходим из горутины
				default: // Иначе (если в chain ничего нет) продолжаем работать
					time.Sleep(1 * time.Second)
					fmt.Println("g[", i, "] is workging")
				}
			}
		}(i)
	}

	// Здесь main горутина ждет сигнала от ОС (Crtl + C)
	s := <-c
	fmt.Println("\nGot signal:", s)
	// как только получили Ctrl+C, закрываем канал chain
	close(chain)
	// ждем пока все воркеры закончат свою работу
	wg.Wait()
}
