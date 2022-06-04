package main

import "fmt"

// Una interfaz es una forma de decir, todos aquellos typos que tengan estos
// metodos como reciever, seran tambien de tipo bot, teniendo acceso a los
// metodos de este pero conservando las propiedades del tipo original
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

/*
func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	smt.Println(sb.getGreeting())
}
*/

// Como podemos ver aunque ambas se llaman igual, tienen implementaciones muy
// diferentes

// Los reciever que solo tienen el tipo es porque no se va a hacer uso del
// valor, esto es perfectamente valido
func (englishBot) getGreeting() string {
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
