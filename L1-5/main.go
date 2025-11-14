package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// Таймаут на канал

// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала – читать эти значения. По истечении N секунд программа должна завершаться.

// Подсказка: используйте time.After или таймер для ограничения времени работы.

func main() {
	c := make(chan int)
	// Это таймер - через 10 секунд в него что-то попадет
	timeAfterChn := time.After(10 * time.Second)
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Горутина которая генерит и отправляет числа
	go func() {
		defer wg.Done()
		for {
			select {
			case _, ok := <-c: // Проверяем не закрыт ли канал
				if !ok {
					fmt.Println("Goroutine is closed")
					return
				}
			default: // Если канал открыт - отправляем число
				number := rand.Int()
				c <- number
				fmt.Println("Goroutine send number:", number)
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// Главный цикл - либо читаем числа, либо ждем таймаута
	for {
		select {
		case m := <-c: // Получили число из канала
			fmt.Println("Main get number:", m)
		case <-timeAfterChn: // Прошло 10 секунд - пора завершаться
			fmt.Println("timed out")
			close(c) // Закрываем канал чтобы горутина поняла что пора останавливаться
			wg.Wait()
			return
		}
	}
}
