package main

import "fmt"

// Группировка температур

// Дана последовательность температурных колебаний:
// -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить эти значения в группы с шагом 10 градусов.
// Пример: -20:{-25.4, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20:{24.5}, 30:{32.5}.
// Пояснение: диапазон -20 включает значения от -20 до -29.9,
// диапазон 10 – от 10 до 19.9, и т.д.
// Порядок в подмножествах не важен.

func main() {
	temperaturesSequence := []float64{
		-25.4, -27.0, 13.0, 13.0,
		19.0, 15.5, 24.5, -21.0,
		32.5, 44.1, -40.3,
		102, -234, -2, -5, -6}
	result := tenGraduceStepJoin(temperaturesSequence)
	for i, v := range result {
		fmt.Println(i, v)
	}
}

func tenGraduceStepJoin(sequence []float64) map[int][]float64 {
	groups := make(map[int][]float64)

	for _, v := range sequence {
		key := int(v/10) * 10
		groups[key] = append(groups[key], v)
	}

	return groups
}
