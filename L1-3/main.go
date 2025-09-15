package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"sync"
	"time"
)

// Работа нескольких воркеров

// Реализовать постоянную запись данных в канал (в главной горутине).
// Реализовать набор из N воркеров, которые читают данные из этого канала и выводят их в stdout.
// Программа должна принимать параметром количество воркеров и при старте создавать указанное число горутин-воркеров.

// постоянный записыватель - канал (главная горутины) - воркеры которые читают из канала - выводят данные в stdout

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Error: Wrong number of args")
		return
	}
	workersNum, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Error: something wrong with args:", err)
		return
	}
	// fmt.Println(workersNum)

	mainChan := make(chan int)
	// mt := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	wg.Add(workersNum)

	// workers
	for i := range workersNum {
		go func() {
			// mt.Lock()
			for {
				num := <-mainChan
				// num, ok := <-mainChan
				// if !ok {
				// 	break
				// }
				fmt.Println("G", i, " |", num)
			}
			// wg.Done()
			// mt.Unlock()
		}()
	}

	// постоянный записватель
	for {
		num := rand.Int()
		time.Sleep(1 * time.Second)
		mainChan <- num
		fmt.Println("main >> ", num)
	}
	// close(mainChan)
	// wg.Wait()
}
