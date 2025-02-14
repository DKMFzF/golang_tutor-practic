package main

import (
	"errors"
	"fmt"
	"math"
)

// композиция интерфейсов
type ShapeWithArea interface {
	Area() float32
}

type ShapeWithPerimetr interface {
	Perimetr() float32
}

type Shape interface {
	ShapeWithArea
	ShapeWithPerimetr
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

func (s *Square) Perimetr() float32 {
	return s.sideLen * 4
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

func (c *Circle) Perimetr() float32 {
	return c.radius * 2 * math.Pi
}

func main() {
	square := newSquare(2.0)
	circle := newCircle(2.0)

	printObject(square)
	printObject(circle)

	printShapeArea(square)
	printShapeArea(circle)

	printShapePerimetr(square)
	printShapePerimetr(circle)

	printObject(square)
	printObject(circle)

	// приведение типов
	fmt.Println(lenString(123))
	var a interface{} = 10
	if m, ok := a.(int); ok {
		fmt.Println(m)
	} else {
		fmt.Println("не удалось привести интерфейс к int")
	}
}

func printShapeArea(s Shape) {
	fmt.Println(s.Area())
}

func printShapePerimetr(s Shape) {
	fmt.Println(s.Perimetr())
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
	if str, ok := i.(string); ok {
		return len(str), nil
	}
	return 0, errors.New("не удалось привести интерфейс к строке")
}
