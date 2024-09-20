package main

import (
	"fmt"
	"os"
)

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
			login(username, password) // Example output
		} else if input == 2 {
			fmt.Println("Invalid input, please try again.")
			fmt.Print("Enter username: ")
			fmt.Scan(&username)
			fmt.Print("Enter password: ")
			fmt.Scan(&password)
			register(username, password)
		} else {
			os.Exit(0)
		}
	}
}
