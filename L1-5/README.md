# L1-5: Таймаут на канал

## Запуск:
```bash
go run main.go
```

## Что делает:
Программа отправляет случайные числа в канал и читает их обратно. Через 10 секунд автоматически завершается.

## Пример вывода:
```
Goroutine send number: 902806090551334903
Main get number: 902806090551334903
Goroutine send number: 8903563171604137321
Main get number: 8903563171604137321
Goroutine send number: 9034727613333465807
Main get number: 9034727613333465807
...
timed out
Goroutine is closed
```