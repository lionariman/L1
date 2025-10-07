package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"runtime"
	"sync"
	"time"
)

// Остановка горутины

// Реализовать все возможные способы остановки выполнения горутины.
// Классические подходы: выход по условию, через канал уведомления, через контекст, прекращение работы runtime.Goexit() и др.
// Продемонстрируйте каждый способ в отдельном фрагменте кода.

// Реализация остановки 10 горутин если используется небуферизированный канал
// гурутины выполняют работу до тех пор пока не получат значение
// в канал notifyChannel который будет являться условием для завершения работы
func stopByChannelSend() {
	wg := sync.WaitGroup{}
	wg.Add(10)

	fmt.Println("Через канал уведомления")

	notifyChannel := make(chan bool)
	defer close(notifyChannel)
	for i := range 10 {
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case <-notifyChannel:
					fmt.Println("Goroutine [", i, "] is closed")
					return
				default:
					fmt.Println("Goroutine [", i, "] is working")
					time.Sleep(1 * time.Second)
				}
			}
		}(i)
	}

	for {
		randNumber := rand.IntN(5)
		fmt.Println("Main goroutine is working | number:", randNumber)
		time.Sleep(2 * time.Second)
		// пусть условием для завершения горутин будет число 3
		if randNumber == 3 {
			fmt.Println("Send TRUE to goroutines to complete them")
			// отправляем 10 раз потому что 10 горутин
			// иначе завершиться только 1 горутина
			// на самом деле я понимаю что это не очень удобно
			// гораздо лучше завершать работу горутин через close
			for range 10 {
				notifyChannel <- true
			}
			break
		}
	}

	wg.Wait()
	fmt.Println("All Goroutines are closed")
}

// Реализация остановки 10 горутин если используется буферизированный канал
// Такой вариант остановки лучше так как горутины не ждут пока канал освободится
// То есть все сообщения сразу попадают в канал и горутины читают из канала по мере готовности
func stopByChannelSend2() {
	wg := sync.WaitGroup{}
	wg.Add(10)

	fmt.Println("Через канал уведомления")

	notifyChannel := make(chan bool, 10)
	defer close(notifyChannel)
	for i := range 10 {
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case <-notifyChannel:
					fmt.Println("Goroutine [", i, "] is closed")
					return
				default:
					fmt.Println("Goroutine [", i, "] is working")
					time.Sleep(1 * time.Second)
				}
			}
		}(i)
	}

	for {
		randNumber := rand.IntN(5)
		fmt.Println("Main goroutine is working | number:", randNumber)
		time.Sleep(2 * time.Second)
		// пусть условием для завершения горутин будет число 3
		if randNumber == 3 {
			fmt.Println("Send TRUE to goroutines to complete them")
			for range 10 {
				notifyChannel <- true
			}
			break
		}
	}

	wg.Wait()
	fmt.Println("All Goroutines are closed")
}

// Остановка горутин через закрытие буферизированного канала
// Эта реализация лучше и проще чем реализация посредством уведомления в канал о закрытии
// так как все горутины мгновенно получают сигнал о завершении
// также при такой реализации меньше кода и меньше используется ресурсов
// и не надо знать количество горутин чтобы закрыть их
func stopByChannelClose() {
	wg := sync.WaitGroup{}
	wg.Add(10)

	fmt.Println("Через канал close(channel)")

	notifyChannel := make(chan bool, 10)
	defer close(notifyChannel)
	for i := range 10 {
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case <-notifyChannel:
					fmt.Println("Goroutine [", i, "] is closed")
					return
				default:
					fmt.Println("Goroutine [", i, "] is working")
					time.Sleep(1 * time.Second)
				}
			}
		}(i)
	}

	for {
		randNumber := rand.IntN(5)
		fmt.Println("Main goroutine is working | number:", randNumber)
		time.Sleep(2 * time.Second)
		// пусть условием для завершения горутин будет число 3
		if randNumber == 3 {
			fmt.Println("Send TRUE to goroutines to complete them")
			// все горутины одновременно получат сигнал о завершении
			close(notifyChannel)
			break
		}
	}

	wg.Wait()
	fmt.Println("All Goroutines are closed")
}

// завершение горутин через контекст с отменой
func stopByContextWithCancel() {
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(10)

	for i := range 10 {
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("G [", i, "] is closed")
					return
				default:
					fmt.Println("G [", i, "] is working")
					time.Sleep(1 * time.Second)
				}
			}
		}(i)
	}

	for {
		randomCancelNumber := rand.IntN(5)
		fmt.Println("Random number", randomCancelNumber)
		// пусть условием для завершения горутин будет число 3
		if randomCancelNumber == 3 {
			fmt.Println("<<< Closing all goroutines >>>")
			cancel()
			break
		}
		fmt.Println("Main G is working")
		time.Sleep(1 * time.Second)
	}

	wg.Wait()
}

// завершение горутин через контекст с дедлайном
func stopByContextWithDeadline() {
	wg := sync.WaitGroup{}

	// дедлайном: завершим работу горутин по истечении 5 секунд
	cancelTime := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), cancelTime)
	defer cancel()

	wg.Add(10)

	for i := range 10 {
		go func(ctx context.Context, i int, wg *sync.WaitGroup) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("G [", i, "] is closed:", ctx.Err())
					return
				default:
					fmt.Println("G [", i, "] is working")
					time.Sleep(1 * time.Second)
				}
			}
		}(ctx, i, &wg)
	}

	wg.Wait()
}

// реализация завершения горутин через контекст с таймаутом
func stopByContextWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := range 10 {
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("G", i, "is closed")
					return
				default:
					fmt.Println("G", i, "is working")
					time.Sleep(1 * time.Second)
				}
			}
		}(i)
	}
	wg.Wait()
}

// runtime.Goexit
// нужен когда требуется мгновенно завершить работу конкретной горутины
// и также надо чтобы были выполнены все defer этой горутины
// это больше "аварийный выход" - нужен при специфических случаях
//
// В примере ниже я запускаю как обычно 10 горутин, но 1 из них
// будет завершен аварийно через Goexit. Остальные горутины
// поработают некоторое время пока не случится таймаут и они все не завершатся
func stopByRuntimeGoexit() {
	wg := sync.WaitGroup{}

	// так как 1 горутина завершается аварийно через runtime.Goexit,
	// то для остальных горутин будет контекст с таймаутом для отложенного завершения
	// думаю, в рамках примера, будет достаточно 5 секунд таймаута
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	wg.Add(10)

	go func() {
		defer wg.Done()
		defer fmt.Println("1 defer")
		defer fmt.Println("2 defer")
		defer fmt.Println("3 defer")
		fmt.Println("G (runtime.Goexit) is working")
		// пусть "поработает" 3 секунды после которого будет вызван Goexit
		time.Sleep(3 * time.Second)
		fmt.Println("G (runtime.Goexit) is closed!")
		runtime.Goexit()
	}()

	for i := range 9 {
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("G", i, "is closed")
					return
				default:
					fmt.Println("G", i, "is working")
					time.Sleep(1 * time.Second)
				}
			}
		}(i)
	}

	wg.Wait()
}

func main() {
	// stopByChannelSend()
	// stopByChannelSend2()
	// stopByChannelClose()
	// stopByContextWithCancel()
	// stopByContextWithDeadline()
	// stopByContextWithTimeout()
	stopByRuntimeGoexit()
}
