package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	/* for i := 0; i < len(links); i++ {
		go checkLink(<-c, c)
	} */

	for {
		go checkLink(<-c, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}

	fmt.Println(link, "is up")
	c <- link
}

// Una go rutine es un proceso de ejecucion de un programa linea a linea, la mas
// famosa es la main go rutine que es la principal de cualquier programa

// Algunas instrucciones en go bloquean la ejecucion normal de un programa, como
// podria ser http.Get por ejemplo

// Una forma de solventar esto es la de crear nuevas routines, es decir, crear
// nuevos entronos de ejecucion de Go a traves de la instruccion go

// Go tiene un organizador de routines llamado scheduler que es el encargado de
// decidir que routine se activa y que routine se pausa

// La forma en la que trabajan las routines difiere en funcion del numero de cpu
// que tenga nuestro ordenador, en caso de que tenga uno, el scheduler solo
// podra ejecutar una routine al mismo tiempo, y solo pasara a ejecutar otra
// cuando esta termine o quede bloqueada (por un codigo como http)

// En caso de que tengamos multiples cpu si que podemos tener un codigo en
// paralelo ya que Go asigna una routine a cada una de las cpu, pero solo una
// cpu por routine nunca mas

// Lo que si ha de quedar muy claro es que no es lo mismo paralelismo que
// concurrencia y Go hace uso de este ultimo, en caso de que el scheduler
// detecte que hay libres varias cpu hara uso de esas cpu al mismo tiempo, pero
// si se da el caso de que cada cpu tenga varias routines por terminar las
// ira cogiendo una a una hasta que se bloqueen o finalicen, por lo que Go es
// concurrente en la practica, no paralelo como tal, tan solo usa el paralelismo
// en parte cuando hacemos uso de muchas cpu a la vez

// Por defecto cuando la main routine crea mas routines este las ignora y
// continua con el proceso de ejecucion de la main routine como si nada, para
// que la main routine este consciente de lo que hacen las otras routines
// tenemos que hacer uso de los channels

// Los channels son canales de comunicacion entre routines tipadas para asegurar
// la estabilidad de lo que transmitimos entre estas

// Cuando queremos insertar informacion en un channel, usamos la sintaxis:
// channel <- valor

// Cuando queramos recibir un valor de un channel tenemos que usar la sintaxis:
// variable <- channel, ten en cuenta que este valor esperara hasta que reciba
// el mensaje, es decir, es bloqueante

// Print puede acceder a los mensajes sin necesidad de guardarlos antes en una
// variable fmt.Println(<-c)

// Otro punto a tener en cuenta de los canales, es que en el momento en que el
// main routine reciba una respuesta de un canal pasara a la siguiente linea de
// comando, por lo que si tenemos un ejemplo como el anterior en que tenemos
// cinco vinculos, deberemos preparar cinco metodos a la espera de recibir una
// respuesta de un canal, sino solo se tendra el cuenta el primero
