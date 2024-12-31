package main

import "fmt"

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	revenue := getUserInput("Enter the revenue: ")
	expenses := getUserInput("Enter the expenses: ")
	taxRate := getUserInput("Enter the tax rate: ")

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	futureValue, ratio := calculateFutureValueAndRatio(ebt, taxRate, profit)

	fmt.Printf("ebt = %.2f\n", ebt)
	fmt.Printf("profit = %.2f\n", profit)
	fmt.Printf("ratio = %.2f\n", ratio)

	// Sprint hace lo mismo que Print, pero devolviendo un string

	// Con el backtick podemos crear multiples lineas, ten en cuenta que se
	// imprimiran tal cual
	taxRateOnString := fmt.Sprintf(`Future value
		%.2f`, taxRate)
	fmt.Println(taxRateOnString)

	outputText(fmt.Sprint(futureValue))
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValueAndRatio(ebt, taxRate, profit float64) (float64, float64) {
	fv := ebt * (1 - taxRate/100)
	r := ebt / profit
	return fv, r
}

func getUserInput(text string) float64 {
	var ui float64
	fmt.Print(text)
	fmt.Scan(&ui)
	return ui
}
