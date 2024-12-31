package main

import (
	"fmt"

	"example.com/investment_calculator/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFileName = "balance.txt"

func main() {
	accBalance, err := fileops.GetFloatFromFile(accountBalanceFileName)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Welcome to Go Bank")
	fmt.Println(randomdata.PhoneNumber())

	for {
		presentOptions()

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your balance is %.2f\n", accBalance)
		case 2:
			fmt.Println("Enter your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount, must be greatter than 0")
				continue
			}

			accBalance += depositAmount
			fmt.Printf("Balance updated, new amount: %.2f\n", accBalance)
			fileops.WriteBalanceToFile(accBalance, accountBalanceFileName)
		case 3:
			fmt.Println("Enter your withdrawal: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount, must be greatter than 0")
				continue
			}

			if withdrawalAmount > accBalance {
				fmt.Println("Sorry, you dont have enough money")
				continue
			}

			accBalance -= withdrawalAmount
			fmt.Printf("Balance updated, new amount: %.2f\n", accBalance)
			fileops.WriteBalanceToFile(accBalance, accountBalanceFileName)
		case 4:
			fmt.Println("Goodbye!")
			// En un switch no podemos usar break, ya que interfiere con la
			// instruccion del bucle for, dentro de un switch, el break es para
			// romper el switch, no el bucle, por lo que tenemos que usar return
			return
		default:
			fmt.Println("Unkown option")
			break
		}

		// if choice == 1 {
		// 	fmt.Printf("Your balance is %.2f\n", accBalance)
		// } else if choice == 2 {
		// 	fmt.Println("Enter your deposit: ")
		// 	var depositAmount float64
		// 	fmt.Scan(&depositAmount)

		// 	if depositAmount <= 0 {
		// 		fmt.Println("Invalid amount, must be greatter than 0")
		// 		continue
		// 	}

		// 	accBalance += depositAmount
		// 	fmt.Printf("Balance updated, new amount: %.2f\n", accBalance)
		// } else if choice == 3 {
		// 	fmt.Println("Enter your withdrawal: ")
		// 	var withdrawalAmount float64
		// 	fmt.Scan(&withdrawalAmount)

		// 	if withdrawalAmount <= 0 {
		// 		fmt.Println("Invalid amount, must be greatter than 0")
		// 		continue
		// 	}

		// 	if withdrawalAmount > accBalance {
		// 		fmt.Println("Sorry, you dont have enough money")
		// 		continue
		// 	}

		// 	accBalance -= withdrawalAmount
		// 	fmt.Printf("Balance updated, new amount: %.2f\n", accBalance)
		// } else {
		// 	fmt.Println("Goodbye!")
		// 	break
		// }
	}

	// fmt.Println("Thanks for choosing our bank!")
}
