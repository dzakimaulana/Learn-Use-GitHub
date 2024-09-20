package main

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	Username string
	Password string
}

var accounts []Account

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func register(username string, password string) error {
	if username == "" || password == "" {
		return errors.New("username or password cannot be empty")
	}

	for _, account := range accounts {
		if account.Username == username {
			return errors.New("username already exists")
		}
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}

	accounts = append(accounts, Account{
		Username: username,
		Password: hashedPassword,
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
			match := checkPasswordHash(password, account.Password)
			if match {
				return nil
			}
			return errors.New("username or password incorrect")
		}
	}
	return errors.New("username or password incorrect")
}
