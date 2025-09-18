package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"time"
)

// У нас есть один начальник который кидает задачи в общую кучу - в канал
// и множество воркеров которые берут задачи из этой кучи и выполняют

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

	// Создаем канал - это как конвейерная лента, по которой передаем данные
	mainChan := make(chan int)

	// Запускаем воркеры - каждый будет читать из канала и печатать
	for i := range workersNum {
		go func() {
			for {
				// Ждем когда в канал что-то положат и берем это
				num := <-mainChan
				fmt.Println("G", i, " |", num)
			}
		}()
	}

	// А тут главная горутина постоянно генерит и кидает данные в канал
	for {
		// Генерируем случайное число
		num := rand.Int()
		// Немножко ждем чтобы не спамить
		time.Sleep(1 * time.Second)
		// Кладем число в канал чтобы один из воркеров его взял в работу
		mainChan <- num
		fmt.Println("main >> ", num)
	}
}
