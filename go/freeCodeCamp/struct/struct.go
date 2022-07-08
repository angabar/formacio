package main

import "fmt"

// Los struct siguen las mismas convenciones que las variables a la hora de
// definir los nombres de las mismas, con la peculiaridad de que si queremos que
// el tipo y sus propiedades sean exportables tenemos que poner el tipo y sus
// propiedades en mayuscula, de lo contrario solo estaremos exportando el tipo
type Doctor struct {
	Number     int
	ActorName  string
	Companions []string
}

type Animal struct {
	Name   string
	Origin string
}

// En Go no tenemos herencia, pero podemos inyectar un struct en otro para que
// este disponga de las propiedades del otro
type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

// Un struct es una estructura de tipos que conforman un objeto, podemos anidar
// esto de las manera que queramos, incluso crear un struct dentro de otro
// struct
func main() {
	// En un struct se puede prescindir de las claves a las que hacen referencia
	// las propiedades, pero no es nada recomendable
	aDoctor := Doctor{
		Number:     3,
		ActorName:  "Perico",
		Companions: []string{"test", "test2"},
	}

	fmt.Println(aDoctor)

	// Para consultar un atributo usamos el punto .
	fmt.Println(aDoctor.Number)

	// Los struct tambien pueden ser anonimos, estos se declaran de la siguiente
	// forma
	person := struct{ name string }{name: "Perico"}
	fmt.Println(person)

	// A diferencia de los map, los struct son pasados o copiados por copia, no
	// por referencia, por lo que modificar uno no alterara el otro, al igual
	// que con los array, si quisieramos modificar tambien el elemento original,
	// tendriamos que guardar o pasar la direccion de memoria, no el elemento en
	// si

	b := Bird{
		Animal: Animal{
			Name:   "jaja",
			Origin: "Idont know",
		},
		SpeedKPH: 23.2,
		CanFly:   true,
	}

	fmt.Println(b.Animal.Name)
	fmt.Println(b.SpeedKPH)
}
