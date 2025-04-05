package main

import "fmt"

func main() {
	// Un slice es un array sin longitud determinada
	prices := []float64{10.00, 9.99}
	fmt.Println(prices[0:1])

	// Podemos asignar valores a indices que existen
	prices[1] = 2.33

	// append añade el elemento a un slice pero devuelve un slice totalmente
	// nuevo, el slice original no se modifica
	updatedPrices := append(prices, 4.55)
	fmt.Println(updatedPrices)
}

// func main() {
// 	// Podemos declarar los arrays de estas dos formas, al inicializarlo como en
// 	// el primer ejemplo no es necesario que los llenemos todos
// 	var productNames [10]string = [10]string{"A book"}
// 	prices := [4]float64{1.0, 2.0, 3.0, 4.0}

// 	productNames[2] = "A carpet"

// 	fmt.Println(prices)
// 	fmt.Println(prices[2])
// 	fmt.Println(productNames)

// 	// Asi es como se corta un array, marcamos la posicion inicial y la final
// 	// (esta ultima no incluida)

// 	// Si omitimos alguno de los dos, se tomara por sentado que se toma el
// 	// principio y el final del array
// 	featuredPrices := prices[1:3]
// 	fmt.Println(featuredPrices)

// 	// length es la longitud del array como tal, cap es la capacidad que tiene
// 	// el array al que estamos apuntando (el slice es un puntero a un array)
// 	fmt.Println(len(featuredPrices), cap(featuredPrices))
// }
