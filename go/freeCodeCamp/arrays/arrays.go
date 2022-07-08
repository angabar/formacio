package main

import "fmt"

func main() {
	// Un array debe estar definido por su tamaño y su tipo, la pega de usar el
	// tipo definido es que estamos declarando el tamano dos veces, cuando lo
	// ponemos entre corchetes y al listar los elementos. Si lo que queremos es
	// que tome el tamaño en funcion de los elementos que tenga, usaremos el ...
	grades := [...]int{10, 11, 23}

	// Podemos definir un array sin valores por defecto y asignarlos despues
	var students [3]string
	students[0] = "Lisa"

	// Podemos averiguar la longitud de un array usando la funcion len()
	fmt.Println(len(students))

	// Podemos tambien generar un array de arrays
	var matrix [3][3]int = [3][3]int{{1, 1, 0}, {1, 0, 0}, {0, 0, 0}}
	fmt.Println(matrix)

	fmt.Printf("Grades: %v %v %v", grades[0], grades[1], grades[2])

	// En algunos lenguajes de programación, cuando copiamos un array de una
	// variable a otra, lo que se hace en realidad es copar la referencia de
	// dicho array, por lo que si lo alteramos estamos alterando el array
	// original. Esto en Go no es asi, al copiar un array se copia este desde 0
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	// Si lo que queremos es modificar ambos arrays a la vez, lo que tenemos que
	// hacer es decirle a Go que lo que va a guardarse en b es la direccion de
	// memoria de a, no a en si, y esto se hace con un puntero b:= &a

	// Los slice son arrays con tamaño indefinido y son similares a los arrays
	// con la diferencia de que no tienen un tamaño definido y en que cuando lo
	// pasamos a una funcion o guardamos en una variable, no se pasa o guarda
	// por copia como los arrays, sino por referencia
	s := []int{2, 3, 4}

	// Al igual que con Python podemos guardar trozos de los slice o arrays en
	// variables
	c := s[:1]
	fmt.Println(c)

	// Con los slice tambien podemos usar la funcion make, que recibe el slice
	// con el tipo, la longitud que tendra, y su capacidad m := make([]int, 3,
	// 100)

	// Podemos agregar elementos a un slice con la funcion append, el cual
	// recibe como primer argumento el slice y los elementos (puede haber mas de
	// uno) a poner
	k := []int{}
	k = append(k, 22)
	fmt.Println(k)

	// Lo que hace Go por detras es crear un nuevo array con mas capacidad para
	// los elementos que pongamos, por eso cuando usemos append tenemos que
	// guardarlo en la misma variable que tenia el slice original

	// Otra cosa que podemos hacer con los slice es concatenarlos con el
	// operador spread
	o := []int{2, 3}
	q := []int{4, 5}
	o = append(o, q...)
	fmt.Println(o)

	// Para eliminar elementos intermedios de un slice, lo que podemos hacer es
	// recortar el slice por el principio y el final
	g := []int{1, 2, 3, 4, 5, 6}
	r := append(g[:2], g[3:]...)
	fmt.Println(r)
}

// Diferencia entre len y cap, un slice usa por debajo un array, len es el
// numero de elementos presentes en el array y cap es la capacidad del array al
// que Go asigna al slice, cuando definimos una capacidad determinada con make
// Go lo que hace es reservarnos un array con n capacidad, si en algun momento
// guardamos mas elementos de lo que habiamos dicho inicialmente Go hara una
// reasignacion de memoria y aumentara su capacidad esto aunque parezca algo
// bueno en realidad afecta al rendimiento por lo que no es algo de lo que
// debamos abusar
/*
Definimos un slice con capacidad de 3
s := make([]int, 0, 3)
for i := 0; i < 5; i++ {
    s = append(s, i)
    fmt.Printf("cap %v, len %v, %p\n", cap(s), len(s), s)
}

Observa como cuando superamos los 3 elementos Go reasigna el array en memoria
cap 3, len 1, 0x1040e130
cap 3, len 2, 0x1040e130
cap 3, len 3, 0x1040e130
cap 6, len 4, 0x10432220
cap 6, len 5, 0x10432220
*/
