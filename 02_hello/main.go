package main

import "fmt"

func main() {
	var name = "Edward"
	var age int64 = 28
	const isCool = true

	fmt.Println(name, age, isCool)

	fmt.Printf("%T\n", age)
}
