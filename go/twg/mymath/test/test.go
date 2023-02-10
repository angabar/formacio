package main

import (
	"fmt"

	"github.com/encece/mymath"
)

func main() {
	sum := mymath.Sum([]int{10, -2, 3})
	if sum != 11 {
		msg := fmt.Sprintf("FAIL: Wanted 11 but received %d", sum)
		panic(msg)
	}

	add := mymath.Add(5, 10)
	if add != 15 {
		msg := fmt.Sprintf("FAIL: Wanted 15 but received %d", sum)
		panic(msg)
	}

	fmt.Println("PASS")
}
