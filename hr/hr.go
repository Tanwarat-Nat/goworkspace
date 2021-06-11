package hr

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	Designation string
}
