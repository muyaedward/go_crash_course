package main

import "fmt"

func main() {
	//Define map
	//emails := make(map[string]string)

	// Assign kv

	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"
	// emails["Mike"] = "mike@gmail.com"

	//Declare map and add key values

	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	// fmt.Println(emails)
	// fmt.Println(len(emails))
	// fmt.Println(emails["Bob"])

	// delete(emails, "Bob")

	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
}
