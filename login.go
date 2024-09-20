package main

import (
	"errors"
	"fmt"
)

type Account struct {
	Username string
	Password string
}

var accounts []Account

func register(username string, password string) error {
	if username == "" || password == "" {
		return errors.New("username or password cannot be empty")
	}

	accounts = append(accounts, Account{
		Username: username,
		Password: password,
	})

	fmt.Printf("Successfully register %s", username)
	return nil
}

func login(username string, password string) error {
	if username == "" || password == "" {
		return errors.New("username or password cannot be empty")
	}

	for _, account := range accounts {
		if account.Username == username {
			if account.Password == password {
				return nil
			}
		}
		return errors.New("username or password incorrect")
	}
	return errors.New("username or password incorrect")
}
