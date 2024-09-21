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
}
