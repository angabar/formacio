package main

import "fmt"

// El primer tipo primitivo de Go es bool para representar booleanos, estos
// pueden guardar valores true o false o expresiones lógicas como 1 == 1 o 3 <
// 5, los bool se inicializan a false por defecto

// El siguiente tipo son los enteros representados por int, en Go tenemos los
// tipos int genericos, los cuales se ponen por defecto si no especificamos el
// tamaño o tipos int de tamaño especifico, cuando queramos especificar el rango
// numerico que va a guardar la variable, tenemos int8, 16, 32, 64

// Enteros sin asignar son 0

// Similar al entero tenemos el entero sin signar, unsigned int, el cual es
// identico al int pero con la diferencia de que no admite valores negativos
// (uint)

// Aunque parezcan del mismo tipo, un int8 no es igual a un int, por lo que no
// podemos usarlos combinados para hacer operaciones matematicas, tendremos que
// convertirlos

// Igual que los enteros tenemos los numeros decimales, los cuales son guardados
// en el tipo float, que igual que con int pueden o no especificar la longitud
// en bits

// Decimales sin asignar valen 0

// Nota: En float no podemos usar el operador remainder (%)

// Otro tipo, numero complejos representados por complex64 y complex128? Tambien
// admite una funcion complex() la cual recibe dos float y devuelve un numero
// complejo

// El valor por defecto de los valores complejos es (0+0i)

// Otro tipo es el string, este es una cadena de texto entre comillas dobles,
// puede ser tratado como una cadena de caracteres y acceder a los elementos de
// uno en uno como si fuese un array (ms[2]) pero teniendo en cuenta que son
// inmutables, no podemos modificar el valor del mismo (ms[2] = "a" error!)

// Otra cosa que podemos hacer con los string es concatenarlos usando + o
// convertir un string en un listado de bytes b := []byte(s)
func main() {
	var n bool = true

	fmt.Println(n)
}
