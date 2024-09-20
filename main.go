package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func helloToYou(person Person) {
	fmt.Printf("Hello %s\n", person.Name)
}

func main() {
	person1 := Person{
		Name: "Dzaki",
		Age:  20,
	}
	helloToYou(person1)
}
