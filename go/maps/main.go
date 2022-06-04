package main

import "fmt"

func main() {
	/* colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	fmt.Println(colors) */

	// var colors map[string]string

	/* colors := make(map[string]string)

	colors["white"] = "#ffffff"

	delete(colors, "white") */

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(mapToPrint map[string]string) {
	for key, value := range mapToPrint {
		fmt.Println(key, value)
	}
}

// diff structs
// map, claves y valores todas del mismo tipo
// map pueden indexarse, acceder con un indice
// map no necesitas saber todos los valores en el momento de la compilacion
// map pasa por referencia, los structs por copia
