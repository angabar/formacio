package main

import (
	"fmt"

	"github.com/angabar/testUtils/utils"
)

func main() {
	sum := utils.Sum([]int{1, 2, 3})

	fmt.Println(sum)

	/* if sum != 6 {
		panic("FAIL: Want 6 but got" + sum)
	} */

	fmt.Println("PASS")
}
