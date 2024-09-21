package main

type Items struct {
	Name  string
	Type  string
	Price int
	Stock int
}

var stuff = []Items{
	{Name: "Aqua", Type: "Dring", Price: 5000, Stock: 10},
	{Name: "Mamoka", Type: "Food", Price: 4000, Stock: 20},
	{Name: "Le Minerale", Type: "Drink", Price: 4000, Stock: 50},
	{Name: "Chitato", Type: "Food", Price: 8000, Stock: 15},
	{Name: "Lays", Type: "Food", Price: 9000, Stock: 15},
}
