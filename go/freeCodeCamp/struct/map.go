package main

import "fmt"

func main() {
	// Los maps son colecciones de pares clave valor, podemos elegir el tipo que
	// queramos con las claves y con los valores, pero una vez definido no
	// podemos cambiarlo
	statePopulations := map[string]int{
		"España":   1,
		"Sueca":    1,
		"Algemesi": 1,
	}

	fmt.Println(statePopulations)

	// Podemos usar make para crear un map en un momento determinado y usarlo
	// mas adelante (en un bucle por ejemplo)
	m := make(map[string]int)
	fmt.Println(m)

	// Podemos acceder a un elemento del map usando su key
	fmt.Println(statePopulations["Sueca"])

	// O agregar un par de key value
	statePopulations["Alzira"] = 22

	// Para eliminar un par clave valor usamos delete
	delete(statePopulations, "Sueca")
	fmt.Println(statePopulations)

	// Los maps no se ordenan como los array, no se garantiza que el orden en el
	// que vuelvan sea siempre el mismo

	// Si consultamos por una clave que no existe Go devolvera el valor por
	// defecto del tipo definido (int, 0) y el estado de si lo a encontrado o no
	// como bool
	pop, ok := statePopulations["oho"]
	// 0, false
	fmt.Println(pop, ok)

	// Podemos obtener la longitud de los map usando la funcion len
	fmt.Println(len(statePopulations))

	// Los maps se pasan por refernecia por lo que si lo copiamos o pasamos y
	// modificamos este, se modificara el elemento original
	c := statePopulations
	delete(c, "Algemesi")
	fmt.Println(statePopulations) // map[Alzira:22 España:1]
	fmt.Println(c)                // map[Alzira:22 España:1]
}
