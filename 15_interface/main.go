package main

import (
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
}

func printShapeArea(s Shape) {
	fmt.Println(s.Area())
}

func printObject(i interface{}) {
	switch t := i.(type) {
	case int:
		fmt.Println("int", t)
	case bool:
		fmt.Println("bool", t)
	case string:
		fmt.Println("string", t)
	default:
		fmt.Println("unknown", t)
	}
}
