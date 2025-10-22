package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	R float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.R * c.R
}

func (c Circle) SetR(r float64) {
	c.R = r
}

func GetArea(s Shape) float64 {
	return s.Area()
}

func main() {
	c := &Circle{
		R: 10,
	}

	fmt.Println(GetArea(c)) // ошибка компиляции? - нет, так как тип Circle реализует интерфейс Shape

	c.SetR(15)

	fmt.Println(GetArea(c)) // ошибка компиляции? - нет, так как тип Circle реализует интерфейс Shape
}
