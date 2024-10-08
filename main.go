package main

import (
	"fmt"
	"os"
)

type Person struct {
	account Account
	wallet  Wallet
}

func main() {
	var input int
	var username string
	var password string

	for {
		fmt.Print("Enter your choice (1 to register): ")
		fmt.Scan(&input)
		if input == 1 {
			fmt.Print("Enter username: ")
			fmt.Scan(&username)
			fmt.Print("Enter password: ")
			fmt.Scan(&password)
			err := login(username, password)
			if err != nil {
				os.Exit(0)
			} // Example output
		} else if input == 2 {
			fmt.Println("Invalid input, please try again.")
			fmt.Print("Enter username: ")
			fmt.Scan(&username)
			fmt.Print("Enter password: ")
			fmt.Scan(&password)
			err := register(username, password)
			if err != nil {
				os.Exit(0)
			}
		} else {
			os.Exit(0)
		}
	}
}
