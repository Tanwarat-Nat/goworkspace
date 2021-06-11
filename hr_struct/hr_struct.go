package main

import (
	"fmt"
	"hr"
)

func main() {
	filicity := hr.Employee{
		Person:      hr.Person{Name: "Filicity", Age: 23},
		Designation: "Programmer",
	}
	fmt.Printf("%+v\n", filicity)
	fmt.Println(filicity.Name)
	fmt.Println(filicity.Age)
}
