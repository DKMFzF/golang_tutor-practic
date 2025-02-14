package main

import (
	"errors"
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
}

// структура Квадрата
type Square struct {
	sideLen float32
	area    float32
}

func (s *Square) Area() float32 {
	s.area = s.sideLen * s.sideLen
	return s.area
}

func newSquare(sideLen float32) *Square {
	return &Square{
		sideLen: sideLen,
	}
}

// структура Круга
type Circle struct {
	radius float32
	area   float32
}

func newCircle(radius float32) *Circle {
	return &Circle{
		radius: radius,
	}
}

func (c *Circle) Area() float32 {
	c.area = c.radius * c.radius * math.Pi
	return c.area
}

func main() {
	square := newSquare(2.0)
	circle := newCircle(2.0)

	printObject(square)
	printObject(circle)

	printShapeArea(square)
	printShapeArea(circle)

	printObject(square)
	printObject(circle)

	fmt.Println(lenString(123))
}

func printShapeArea(s Shape) {
	fmt.Println(s.Area())
}

func printObject(i interface{}) {
	// конструкция type switch
	// конструкция которая берёт значения из интерфейса
	// и передаёт её в переменную value
	switch value := i.(type) {
	case int:
		fmt.Println("int", value)
	case bool:
		fmt.Println("bool", value)
	case string:
		fmt.Println("string", value)
	default:
		fmt.Println("unknown", value)
	}
}

// функция которая привеодит интерфейс к строку
func lenString(i interface{}) (int, error) {
	str, ok := i.(string)
	if ok {
		return len(str), nil
	}
	return 0, errors.New("не удалось привести интерфейс к строке")
}
