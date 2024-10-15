package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	height, width float64
}

func (r *Rectangle) Area() float64 {
	return r.height * r.width
}

type Circle struct {
	radius float64
}

// Считает в метрах, а нам нужно в сантиметрах
func (c *Circle) AreaInMeters() float64 {
	return math.Pi * c.radius / 100 * c.radius / 100
}

type CircleAdapter struct {
	*Circle
}

func (c *CircleAdapter) Area() float64 {
	return c.AreaInMeters() * 10000
}

func CalculateArea(shape Shape) float64 {
	return shape.Area()
}

func main() {
	circle := Circle{radius: 3}
	rectangle := Rectangle{height: 5, width: 5}
	adapter := CircleAdapter{&circle}

	fmt.Printf("rectangle area: %f\n", rectangle.Area())
	fmt.Printf("circle area: %f", adapter.Area())
}
