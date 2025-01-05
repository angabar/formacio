package main

import "fmt"

func main() {
	age := 32

	// var agePointer *int --> seria lo mismo
	// agePointer := &age

	// fmt.Printf("Adult years: %d\n", editAgeToAdultYears(agePointer))

	editAgeToAdultYears(&age)

	fmt.Printf("%d\n", age)
}

func editAgeToAdultYears(age *int) {
	// return *age - 18

	*age -= 18
}
