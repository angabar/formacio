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

	jim.updateName("jajajaa")
	jim.print()
}

func (singlePerson person) updateName(newFirstName string) {
	singlePerson.firstName = newFirstName
}

func (singlePerson person) print() {
	fmt.Printf("%+v", singlePerson)
}
