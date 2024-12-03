package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount, years float64 = 10000, 5
	expectedReturnRate := 5.5

	var futureValue = investmentAmount * math.Pow((1+expectedReturnRate)/100, years)
	fmt.Println(futureValue)
}
