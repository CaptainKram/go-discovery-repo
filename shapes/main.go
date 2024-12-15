package main

import "fmt"

func main() {
	t := triangle{base: 4, height: 2}
	sq := square{side: 3}

	printArea(t)
	printArea(sq)
}

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	side float64
}

func (t triangle) getArea() float64 {
	return 0.5 * float64(t.base) * float64(t.height)
}

func (s square) getArea() float64 {
	return float64(s.side) * float64(s.side)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
