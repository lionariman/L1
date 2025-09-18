# L1-4: Graceful Shutdown с воркерами

## Запуск:
```bash
go run main.go <количество_воркеров>
```

### Примеры:
```bash
go run main.go 3    # Запуск с 3 воркерами
go run main.go 5    # Запуск с 5 воркерами
```

## Как остановить:
`Ctrl+C` для корректного завершения всех воркеров

## Пример вывода:
```
Number of Workers: 3
g[ 0 ] is workging
g[ 1 ] is workging
g[ 2 ] is workging
^C
Got signal: interrupt
g[ 0 ] is closed
g[ 1 ] is closed
g[ 2 ] is closed
```