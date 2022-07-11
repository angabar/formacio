package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Strat")
	// defer nos permite ejecutar una expresión en el momento en que se termine
	// la ejecución del bloque en el que se encuentre dicha expresión,
	// independientemente de su posición en dicho bloque
	defer fmt.Println("Middle")

	// En caso de haber varios defer el orden de ejecución es "el último en
	// entrar es el primero en salir", en este caso End entró el último así que
	// en ejecución se ejecutará antes, al final es como la inversa del código
	// de ejecución normal
	defer fmt.Println("End")

	res, err := http.Get("http://www.google.com/robots.txt")

	if err != nil {
		log.Fatal(err)
	}

	// Un caso practico del defer es para cerrar funciones asíncronas que puedan
	// quedar abiertas en el momento de finalizar las acciones que queramos, una
	// petición http, un archivo abierto...

	// Es una buena practica cerrar estos con un defer a continuación de la
	// apertura del mismo
	defer res.Body.Close()

	robots, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", robots)

	// El defer tiene en cuenta los valores en el momento de su ejecución, es
	// decir, si usa una variable que en el momento de llegar al defer vale "x"
	// por mucho que esta variable cambie a lo largo del programa, eso es lo que
	// valdrá cuando finalice la ejecución de la función
	a := "Start"
	defer fmt.Println(a)
	a = "end"

	// En Go no hay excepciones, hay "panics" que son usados en situaciones
	// especiales donde el programa pueda fallar
	fmt.Println("Start")

	empName := "Samia"
	age := 75

	employee(&empName, age)

	// Cuando usamos un panic, todo lo que venga a continuación del mismo no se
	// ejecutará
	/* panic("Error!")
	fmt.Println("End") */

	// Se usa en situaciones donde el objeto err puede no ser nil
	/* if err != nil {
		panic(err.Error())
	} */

	// En caso de que tengamos algun defer antes del panic este se ejecutará
	// antes de salir del programa ya que salir del programa es salir de la
	// función como tal
}

func panicHandler() {
	// La función recover nos permite recuperar el control del programa cuando
	// este a pasado por un panic
	rec := recover()
	if rec != nil {
		fmt.Println("RECOVER", rec)
	}
}

func employee(name *string, age int) {
	defer panicHandler()
	if age > 65 {
		panic("Age cannot be greater than retirement age")
	}

}
