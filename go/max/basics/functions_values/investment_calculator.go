package main

import (
	"fmt"
	"math"
)

func not_main() {
	const inflationRate = 2.5
	var investmentAmount, years float64 = 1000, 5
	var expectedReturnRate float64

	fmt.Print("Enter the invesment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate)/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureRealValue)
}
