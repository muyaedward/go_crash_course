package main

import "fmt"

func main() {
	//Point to memory address of a location of a value
	a := 5
	b := &a //Assign b to a pointer of a

	fmt.Println(a, b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b) //Star represent a pointer

	//Use * to read val from address
	fmt.Println(*b)
	fmt.Println(*&a)

	//Change val with pointer

	*b = 10
	fmt.Println(a)
}
