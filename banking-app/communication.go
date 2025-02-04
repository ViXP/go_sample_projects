package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

func presentOptions() {
	fmt.Println("\nWelcome to OurBank!")
	fmt.Println("Choose your operation: ")
	fmt.Println("1. Check your balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
	fmt.Println("or reach us by the number:", randomdata.PhoneNumber())
}
