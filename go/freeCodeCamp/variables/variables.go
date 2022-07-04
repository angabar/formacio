package main

import "fmt"

// Todas las varibales en go tienen que usarse

// Hay tres tipos de declaracion de variables
// 1. A nivel de package en minuscula, pertenecen al paquete actual
// 2. A nivel de package en mayuscula, son exportadas a otros paquetes
// 3. A nivel de bloque en minuscula, pertenecen al bloque en el que estan

// La longitud del nombre de la variable es importante, un nombre corto
// significa que va a usarse enseguida un nombre largo es que va a ser usada por
// un largo periodo de tiempo

// Los acronimos tenemos que ponerlos en mayuscula siempre que sea posible
// (theURL, theHTTPResquest)

// Cuando declaramos una variable fuera de la funcion main, tiene que ser
// declarada con var si o si
var num int = 22

// Tambien podemos hacer una declaracion de variables en bloque usando var()
var (
	num1 int = 1
	num2 int = 2
)

var global int = 22

func main() {
	// Perfecto si queremos declarar una variable que no vamos a usar aun
	var i int
	i = 42

	// El segundo tipo podriamos decir que es interesante para cuando queremos
	// tener informacion de la variable y del tipo de la misma

	// var i int = 42
	// i := 42

	// Podemos re-declarar global porque una esta fuera del bloque main y la
	// otra dentro, es decir, estan en contextos diferentes, pero si intentamos
	// acceder antes a esta, obtendremos el valor de fuera no el de dentro
	fmt.Println(global)
	var global int = 12

	fmt.Println(global)

	fmt.Println(i)

	var k int = 22
	fmt.Printf("%v, %T\n", k, k)

	var j float64

	// La forma de hacer una conversion de tipos es la siguiente, teniendo en
	// cuenta que se puede producir perdida de informacion
	j = float64(i)
	fmt.Println(j)
}
