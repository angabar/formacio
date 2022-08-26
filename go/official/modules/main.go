package main

import (
	"fmt"
	"log"

	"example.com/greeting/greeting"
)

func main() {
	res, err := greeting.Hello("Perico")

	if err != nil {
		log.Fatalf("Error! %v", err)
	}

	fmt.Println(res)

	ress, err := greeting.Hellos([]string{"Perico", "Paco"})

	if err != nil {
		log.Fatalf("Error! %v", err)
	}

	fmt.Println(ress)
}
