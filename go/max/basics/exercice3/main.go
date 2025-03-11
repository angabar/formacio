package main

import "fmt"

type Product struct {
	id    int
	title string
	price float64
}

func main() {
	hobbies := [3]string{"chess", "running", "cycling"}
	fmt.Println(hobbies)

	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	// La capacidad de este slice es 3, aunque tenga solo dos elementos, porque
	// apunta a un array de 3 elementos
	fmt.Println(cap(mainHobbies))
	// Dado el anterior punto, podemos "expandir" un slice, haciendo que tenga
	// más elementos que los que tenía originalmente
	mainHobbies = mainHobbies[1:3]

	// Mostrará "cycling" también, aunque originalmente no estaba en el array
	fmt.Println(mainHobbies)

	courseGoals := []string{"Learn Go!", "Learn all the basics"}
	fmt.Println(courseGoals)

	courseGoals[1] = "Learn all the details"
	courseGoals = append(courseGoals, "Learn all the basics")
	fmt.Println(courseGoals)

	products := []Product{
		{
			id:    1,
			title: "First product",
			price: 10.99,
		},
		{
			id:    2,
			title: "Second product",
			price: 12.99,
		},
	}

	fmt.Println(products)

	newProduct := Product{
		id:    3,
		title: "Third product",
		price: 3.99,
	}

	products = append(products, newProduct)
	fmt.Println(products)
}
