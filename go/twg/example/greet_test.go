package example

import "fmt"

func ExampleHello() {
	greeting, err := Hello("John")
	if err != nil {
		panic(err)
	}

	fmt.Println(greeting)

	// Output: Hello, John
}

func ExampleDemo_Hello() {
	fmt.Println("Hello")

	// Output: Hello
}

// Los ejemplos sirven para documentar las funciones que escribimos, una de las
// ventajas de Go es que permiten que estos estén testeados, es decir, que los
// ejemplos estén vinculados a tests. De este modo nos aseguramos que siempre
// van a funcionar.

// Para ello escribimos el test sin usar la librería de testing y empezando el
// nombre con Example en lugar de con Test, el Output que aparece justo al final
// del bloque de ejemplo es la aserción que hacemos aunque esté en código
// comentado.

// Si el resultado de la función no es idéntico a lo que pasamos como Output,
// este fallará.

// godoc -http=:6060 para ver la documentación oficial con los ejemplos

// Poniendo una descripción justo encima del nombre del paquete, especificamos
// una pequeña descripción del mismo

// Es importante seguir las reglas de los nombres, especificando todas las
// variedades posibles, si lanzamos godoc con estos ejemplos:

/*

func ExampleHello() {
	...

	// Output: Hello, John
}

func ExampleHello_spanish() {
	...

	// Output: Hola, John
}

*/

// godoc ya nos generará un ejemplo con Hello y otro para Hello (spanish), esto
// mismo se aplica a la hora de diferencias métodos de funciones:

/*

type Demo struct{}

// Hello method
func (d Demo) Hello() {
	...
}

// Hello function
func Hello(name string) (string, error) {
	...
}

func ExampleHello() {
	...

	// Output: Hello, John
}

func ExampleDemo_Hello() {
	...

	// Output: Hello
}

*/

// Es importante seguir una buena nomenclatura de nombres porque de este modo
// godoc nos generará unos documentos fáciles de seguir
