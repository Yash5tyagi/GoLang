package main

import "fmt"

type shape interface {
	CalcArea() float64
}

type square struct {
	side float64
}
type triangle struct {
	height, base float64
}

func (sq square) CalcArea() float64 {
	return sq.side * sq.side
}
func (tri triangle) CalcArea() float64 {
	return tri.base * tri.height / 2
}
func PrintArea(s shape) {
	fmt.Println("Area=", s.CalcArea())
}
func main() {
	s := square{1.2}
	t := triangle{base: 1.1, height: 1.0}
	PrintArea(s)
	PrintArea(t)
}
