package main

import "fmt"

// Cuando especificamos un unico valor a un bloque de variables Go lo que hace
// es aplicar dicho valor al resto de variables que componen el bloque
const (
	a = 2
	b
	c
)

// Otra forma de usar estos bloques es usando la palabra iota, la cual nos
// permite crear un incremento numerico de enteros
const (
	z = iota
	w
	t
)

func main() {
	// Las constantes tienen los mismos requisitos que vimos en las variables en
	// cuanto a nombres, exportaciones, etc
	const myConst int = 42

	// Las constantes no se pueden reasignar myConst = 22

	// Las constantes no pueden guardar datos que deban compilarse, pueden
	// guardar datos que dependan de otros, pero no datos a calcular

	// esto si
	// const seconds = 120
	// const minutes = seconds / 60

	// esto no
	// func addMinutes(minutes int) {
	// 	const more = minutes + 60
	// 	return more
	// }

	// Los tipos primitivos funcionan bien con constantes asi como el shadowing

	// Las constantes, como las variables pueden definirse sin tipo especifico y
	// de este modo operar con otras variables de tipos especificos sin
	// problema, esto tambien pasa cuando son variables las dos
	// var a = 2
	// var a int16 = 3

	fmt.Printf("%v, %T\n", myConst, myConst)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
