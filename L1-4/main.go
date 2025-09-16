package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"time"
)

// Завершение по Ctrl+C

// Программа должна корректно завершаться по нажатию Ctrl+C (SIGINT).
// Выберите и обоснуйте способ завершения работы всех горутин-воркеров при получении сигнала прерывания.
// Подсказка: можно использовать контекст (context.Context) или канал для оповещения о завершении.

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
	chn := make(chan int)
	for i := range numOfWorkers {
		go func() {
			for {
				num := <-chn
				fmt.Println("g[", i, "] > ", num)
			}
		}()
	}
	for {
		randNum := rand.Int()
		time.Sleep(1 * time.Second)
		fmt.Println("main g > ", randNum)
		chn <- randNum
	}
}
