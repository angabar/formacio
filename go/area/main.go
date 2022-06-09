package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}

type square struct {
	base   float64
	height float64
}

type shape interface {
	getArea() float64
}

func main() {
	myTriangle := triangle{
		base:   3,
		height: 6,
	}

	mySquare := square{
		base:   3,
		height: 3,
	}

	printArea(myTriangle)
	printArea(mySquare)
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

func (tr triangle) getArea() float64 {
	return tr.base * tr.height / 2
}

func (sq square) getArea() float64 {
	return sq.base * sq.height
}
