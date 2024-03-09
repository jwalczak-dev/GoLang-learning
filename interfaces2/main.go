package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type square struct {
	sideLength float64
}

type circle struct {
	radius int
	col    string
}

func main() {
	fmt.Println(math.Pi * float64(3))
	circle1 := circle{radius: 4}
	fmt.Println(circle1.area())
	printArea(circle1)

	square1 := square{sideLength: 5}
	printArea(square1)

}

func (s square) area() float64 {
	return s.sideLength * s.sideLength
}

func (c circle) area() float64 {
	return math.Pi * math.Sqrt(float64(c.radius))
}

func printArea(s shape) {
	fmt.Println(s.area())
}
