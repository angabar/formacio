package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Cuando la funcion main termina, todas las rutinas se cierran
func main() {
	c := make(chan int)

	for i := 0; i < 5; i++ {
		go sleepy(i, c)
	}

	for i := 0; i < 5; i++ {
		select {
		case r := <-c:
			fmt.Printf("%v has finished\n", r)
		case <-time.After(2 * time.Second):
			fmt.Printf("Timed out!\n")
		}
	}
}

func sleepy(i int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	c <- i
}
