package main

import (
	"fmt"

	"example.com/banking-app/io"
)

func main() {
	var balance, error = io.ReadFloatFromFile()

	if error != nil {
		fmt.Println(error)
	}

	for {
		presentOptions()
		var usersChoice int
		fmt.Scan(&usersChoice)
		fmt.Println("Your choice: ", usersChoice)

		switch usersChoice {
		case 1:
			fmt.Printf("Your account balance is $%.2f", balance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount > 0 {
				balance += depositAmount
				io.WriteFloatToFile(balance)
				fmt.Printf("Your account balance is $%.2f", balance)
			} else {
				fmt.Println("Invalid amount!")
			}
		case 3:
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if balance >= withdrawAmount {
				balance -= withdrawAmount
				io.WriteFloatToFile(balance)
				fmt.Printf("Your account balance is $%.2f", balance)
			} else {
				fmt.Println("You can't withdraw this sum!")
			}
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing OurBank!")
			io.WriteFloatToFile(balance)
			return
		}
	}
}
