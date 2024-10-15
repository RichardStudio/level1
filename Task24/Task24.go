package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func Distance(p1, p2 *Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

func main() {
	point1 := NewPoint(3, 4)
	point2 := NewPoint(0, 0)

	distance := Distance(point1, point2)
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
