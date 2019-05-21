package main

import (
	"fmt"
	"strconv"
)

//Define person struct

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Types of go methods

// Value receivers
// Pointer receiver

//Greeting method (value reciever)

func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Atoi(p.age)
}

func main() {
	//Init person using struct
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}

	//person1 := Person{"Samantha", "Smith", "Boston", "f", 25}

	//Alternative
	//fmt.Println(person1)

	//fmt.Println(person1.firstName)

	fmt.Println(person1.greet())
}
