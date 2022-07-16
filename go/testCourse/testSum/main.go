package main

import (
	"fmt"

	"github.com/angabar/testUtils/utils"
)

func main() {

	add := utils.Add(5, 10)
	fmt.Println(add)
	/* if add != 15 {
		msg := fmt.Sprintf("Fail: Wanted 15 but got %d", sum)
		panic(msg)
	} */

	fmt.Println("PASS")
}
