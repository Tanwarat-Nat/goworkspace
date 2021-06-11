package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	person      Person
	designation string
}

func main() {
	filicity := Employee{
		person:      Person{"Filicity", 23},
		designation: "Programmer",
	}
	fmt.Printf("%+v\n", filicity)
	fmt.Println(filicity.person.name)
}

/*
การ embeded

 Employee (

	Person (

		name,
		age,
	),

	destination

)

*/
