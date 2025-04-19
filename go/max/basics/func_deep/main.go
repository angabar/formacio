package main

import "fmt"

func main() {
	fac := factorial(5)

	fmt.Println(fac)
}

func factorial(number int) int {
	resutl := 1

	for i := 1; i <= number; i++ {
		resutl = resutl * i
	}

	return resutl
}

recursive
2:56