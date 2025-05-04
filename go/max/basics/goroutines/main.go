package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true

	// Con close cerramos el canal, pero esto solo lo podemos hacer como (en
	// este caso) sabemos cual es la llamada que mas va a tardar, de lo
	// contrario podemos cerrarlo antes de tiempo y quedarnos sin la posibilidad
	// de seguir usando el canal
	close(doneChan)
}

func main() {
	// Es importante entender que si pasamos el mismo canal a todas las
	// funciones, la primera que termine hara terminar el programa y el resto de
	// funciones no se ejecutara
	done := make(chan bool)

	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're liking the course!", done)

	// El programa esperara hasta que todos los canales abiertos comuniquen para
	// cerrar
	// <-done

	for range done {
	}
}
