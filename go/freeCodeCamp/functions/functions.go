package main

import "fmt"

type greeter struct {
	greeting string
	name     string
}

func main() {
	sayMessage("jajaja")
	sayGreeting("hey", "tu")

	name := "Stacey"
	changeName(&name)
	fmt.Println(name)

	sum(1, 2, 3, 4)

	hola := sayHola("jajajaja")
	fmt.Println(hola)

	res := sum2(1, 2)
	fmt.Println(*res)

	res2 := sum3(2, 2)
	fmt.Println(res2)

	res3, err := divide(2.0, 3.0)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res3)

	// Esto es una función anónima, llamada así porque no tiene nombre. Esta
	// concrétamente además es una función inmediata, ya que al tener los
	// paréntesis a continuación la estamos llamando
	func() {
		fmt.Println("juas")
	}()

	// No son muy usadas, pero en el caso de que las necesitemos, tenemos que
	// procurar usar siempre variables del scope de la función anónima, no las
	// de fuera de esta, para así evitar comportamientos extraños
	for i := 0; i < 5; i++ {
		func(p int) {
			fmt.Println(p)
		}(i)
	}

	// Podmeos guardar funciones dentro de una variable, ya sea usando var o no,
	// la ventaja de esto es que podemos definir las plantillas en la
	// declaración y después guardar funciones anónimas que cumplan esos
	// criterios
	var f func() = func() { fmt.Println("juas") }
	f()

	var divide2 func(float64, float64) (float64, error)
	divide2 = func(a, b float64) (float64, error) {
		return a / b, nil
	}
	resw, error := divide2(3.0, 4.0)
	fmt.Println(error)
	fmt.Println(resw)

	myGreet := greeter{
		greeting: "jaja",
		name:     "test",
	}
	myGreet.greetItem()
}

func sayMessage(msg string) {
	fmt.Println(msg)
}

// Si todos los argumentos son del mismo tipo, podemos agruparlos separandolos
// por comas
func sayGreeting(greeting, message string) {
	fmt.Println(greeting, message)
}

// Las funciones pueden recibir punteros para trabajar de una forma más
// eficiente y no re-copiar las variables
func changeName(name *string) {
	*name = "Ted"
	fmt.Println(*name)
}

// Podemos definir un número variable de parámetros con los puntos suspensivos
// en el tipo, solo puede haber uno por función y tiene que estar al final
func sum(values ...int) {
	fmt.Println(values)
}

// Los tipos devueltos se especifican en la parte derecha de la función
func sayHola(msg string) string {
	return msg
}

// Una función también puede devolver un puntero
func sum2(num1, num2 int) *int {
	result := num1 + num2
	return &result
}

// También podemos tener un resturn explícito, esto es definir una variable que
// devolverá la función, con esto hacemos que esta sea utilizable dentro de la
// misma como si ya estuviera declarada. Además de poder prescindir de la misma
// en el return

// No es una técnica muy recomendable para funciones largas
func sum3(num1, num2 int) (result int) {
	result = num1 + num2
	return
}

// Go también permite que las funciones devuelvan múltiples resultados, esto es
// muy útil para cuando vamos a devolver objetos de tipo error. El orden de
// returns será el orden en el que se guardarán las variables
func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("cant divide by 0")
	}

	return a / b, nil
}

// Esto es un método, nos permiten definir funciones que serán utilizables solo
// por los tipos struct que sean del mismo tipo

// Para que una función sea un método, tenemos que definir su reciever, el cual
// establece el tipo de struct que ha de usarse

// Como es un struct, por defecto no es modificable dentro del método, ya que
// estos se pasan por copia, si lo que queremos es modificarlos lo que tenemos
// que hacer es pasar un puntero que apunte al mismo
func (g greeter) greetItem() {
	fmt.Println(g.greeting, g.name)
}
