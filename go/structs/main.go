package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

/* type person struct {
	firstName string
	lastName  string
	contact   contactInfo
} */

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	/* alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex) */

	/* var alex person

	alex.firstName = "test"

	fmt.Println(alex)
	fmt.Printf("%+v", alex) */

	/* jim := person{
		firstName: "Jim",
		lastName:  "Jaja",
		contact: contactInfo{
			email:   "test@test.com",
			zipCode: 123123,
		},
	} */

	jim := person{
		firstName: "Jim",
		lastName:  "Jaja",
		contactInfo: contactInfo{
			email:   "test@test.com",
			zipCode: 123123,
		},
	}

	// jimPointer := &jim
	jim.updateName("jajajaa")
	jim.print()

	mySlice := []string{"Hi", "there", "how", "are", "you"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func (pointerPerson *person) updateName(newFirstName string) {
	(*pointerPerson).firstName = newFirstName
}

func (singlePerson person) print() {
	fmt.Printf("%+v", singlePerson)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}

// Por defecto, el reciever crea una copia del valor pasado, no modifica la
// variable original

// Turn address into values with *address Turn values into
// address with &value

// Si le pasamos a un reciever pointer un valor normal, Go pilla que es una
// referencia no un valor como tal, y lo convierte en una referencia
// automaticamente

// Los slices si que son pasados por referencia a las funciones, a diferencia de
// todos los demas, porque lo que hace Go es crear un array y una estructura de
// control, lo que copia Go es esta estructura como tal, el contenido del array
// sigue siendo el original

// Los otros elementos que se comportan igual son, los maps, los channels, los
// pointers y las funciones
