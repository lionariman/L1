package main

import (
	"fmt"
	"math"
)
	
// Расстояние между точками

// Разработать программу нахождения расстояния
// между двумя точками на плоскости.
// Точки представлены в виде структуры Point
// с инкапсулированными (приватными) полями
// x, y (типа float64) и конструктором.
// Расстояние рассчитывается по формуле между координатами двух точек.

// Подсказка:
// используйте функцию-конструктор NewPoint(x, y),
// Point и метод Distance(other Point) float64.

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func (this *Point) Distance(other Point) float64 {
	return math.Sqrt(math.Pow(other.x - this.x, 2) + math.Pow(other.y - this.y, 2))
}

func main() {
	p1 := NewPoint(2.0, 3.0)
	p2 := NewPoint(-4.0, -3.0)
	rs := p1.Distance(p2)
	fmt.Println(rs)
}