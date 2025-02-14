package shape

import "math"

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
}

func (s *Square) Area() float32 {
	return s.sideLen * s.sideLen
}

func (s *Square) Perimetr() float32 {
	return s.sideLen * 4
}

func NewSquare(sideLen float32) Square {
	return Square{
		sideLen: sideLen,
	}
}

// структура Круга
type Circle struct {
	radius float32
}

func NewCircle(radius float32) Circle {
	return Circle{
		radius: radius,
	}
}

func (c *Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func (c *Circle) Perimetr() float32 {
	return c.radius * 2 * math.Pi
}
