package main

import "fmt"

type Person struct {
	firstName string
	age       int
}

func main() {
	var mike Person
	mike.firstName = "Kohta"
	mike.age = 20
	fmt.Println(mike.firstName, mike.age)
}
