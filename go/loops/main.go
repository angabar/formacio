package main

import (
	"fmt"
	"strings"
)

func main() {
	var isInText = strings.Contains("my loop", "loop")
	fmt.Printf("The condition is: %v\n", isInText)

	var isAppleGreaterThanBannana = "apple" > "banana"
	fmt.Printf("apple is greater than a banana %v\n", isAppleGreaterThanBannana)

	fmt.Printf("2100 is a lap year? %v\n", isLapYear(2100))

	var dayOfMonth = 3

	switch dayOfMonth {
	case 1:
		fmt.Println("Día 1")
	case 2:
		fmt.Println("Día 2")
	case 3:
		fmt.Println("Día 3")
		// En Go podemos usar fallthrough si queremos que se ejecute la
		// sentencia de abajo, por defecto esto no ocurre
		fallthrough
	case 4:
		fmt.Println("Día 4")
	case 22:
		fmt.Println("Día 22")
	case 23, 24, 25:
		fmt.Println("Día 23, 24 o 25")
	default:
		fmt.Println("Default statement")
	}

	var count = 0

	for count < 10 {
		fmt.Println(count)
		count++
	}

	// Los fors pueden o no tener condiciones, la única cosa a tener en cuenta
	// es que sin condición tenemos que pararlo nosotros con un break
	count = 0

	for {
		fmt.Println(count)
		count++

		if count > 10 {
			break
		}
	}
}

func isLapYear(year int) bool {
	// Cuando una operacion booleana usa && han de comprobarse todas las
	// variables para determinar si es true o false (a no ser que la primera ya
	// sea false) sin embargo cuando usamos || si el primer operador ya es true
	// la expresión entera es true por lo que ahorramos tiempo de lógica si
	// ponemos antes el || que el &&
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}
