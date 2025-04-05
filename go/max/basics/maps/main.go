package main

import "fmt"

// Así es como determinamos tipos customizados, para aclaraciones en el código y
// tener este más entendible
type floatMap map[string]float64

// Un tipo custom también puede tener métodos
func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// make nos permite crear slices, maps o channels obteniendo directamente la
	// dirección de memoria del elemento creado, en el caso de los slice nos
	// permite también determinar el tamaño inicual y su capacidad antes de
	// tener que hacer un re-allocate
	userNames := make([]string, 2, 5)

	userNames[0] = "Antonio"

	userNames = append(userNames, "Max")

	fmt.Println(userNames)

	// Con make para los maps tenemos lo mismo, sino usamos make, Go reubica la
	// memoria cada vez que mutamos el map, si usamos make y determinamos el
	// tamaño máximo de este antes de reubicar, estamos haciendo uso de estos de
	// manera más eficiente
	courses := make(floatMap, 3)

	courses["Perico"] = 4.8
	courses["test"] = 33.2

	courses.output()
}
