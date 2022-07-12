package main

import "fmt"

type myStruc struct {
	foo int
}

func main() {
	// Un puntero es una dirección de memoria de otro elemento al que podremos
	// acceder para consultarlo o modificarlo
	var a int = 42

	// Cuando queremos obtener una dirección de memoria usamos & y cuando
	// queremos obtener el contenido de una dirección de memoria usamos *

	// Si el * esta en una variable es que queremos obtener el valor de ese
	// puntero, si esta en un tipo es para definir que ese valor será de tipo
	// puntero
	/*
		var a int = 22
		var b *int = &a
		fmt.Println(*b)
	*/

	var b *int = &a
	fmt.Println(a, b)

	// Cambiemos el valor de a o de b, en realidad estamos cambiando el mismo
	// valor en la misma ubicación
	a = 27
	fmt.Println(a, *b)

	*b = 12
	fmt.Println(a, *b)

	f := [3]int{1, 2, 3}
	e := &f[0]
	q := &f[1]
	fmt.Printf("%v %p %p", f, e, q)

	// Podemos crear un puntero directamente, sin tener que inicializar primero
	// una variable

	// Al hacerlo, si después queremos agregarle un valor, tenemos que usar el
	// operador &

	// Todos aquellos punteros que no sean inicializados tienen el valor de nil
	var ms *myStruc
	ms = &myStruc{foo: 42}
	fmt.Println(ms)
	renderMyStructFoo(ms)

	// Para poder obtener una dirección de memoria, el elemento al que apuntamos
	// a de ser direccioanble, por ejemplo, un struct, o una variable ya
	// declarada

	// Otros elementos como el resultado de una función, un map, un valor
	// determinado no lo son, ya que para estos Go no ha determinado una
	// dirección de memoria en el momento en que hacemos la consulta (en el caso
	// de los map si, pero en este caso es porque los map no guardan un índice
	// determinado, sino que va cambiando).
	/*
		fooPtr := &foo()

		m := map[int64]int64{
		    1: 2,
		}
		a := &m[1]

		i := &4
	*/

	// Los arrays y structs se pasan por copia, los maps y slices por referencia
}

func renderMyStructFoo(ms *myStruc) {
	// No es necesario poner (*ms).foo ya que el compilador entiende que
	// queremos acceder a la propiedad de un puntero struct
	fmt.Println(ms.foo)
}
