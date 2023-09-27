package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var (
		spaceXMaxSpeed              = 100800
		distanceBetweenEarthAndMars = 96300000
	)
	var hoursPerDay, minutsPerHour = 24, 60
	var weight = 60

	fmt.Println("El peso en la superficie de Marte es", 75*0.3783, "y tu edad es", 36*365/687)
	fmt.Printf("Este texto permite poner valores dinámicos %8v\n", 35)

	fmt.Printf("El tiempo que costará llegar es de %v días\n", distanceBetweenEarthAndMars/spaceXMaxSpeed/24)

	fmt.Printf("Hay %v horas en un día y %v minutos en una hora\n", hoursPerDay, minutsPerHour)

	weight *= 2
	fmt.Printf("El peso duplicado es %v\n", weight)

	var randNum = rand.Intn(10) + 1
	fmt.Printf("First value: %v\n", randNum)
	randNum = rand.Intn(10) + 1
	fmt.Printf("First value: %v\n", randNum)
}
