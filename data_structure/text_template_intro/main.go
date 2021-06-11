package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	filicity := Person{"Filicity", 23}
	oliver := Person{"Oliver", 25}

	const greetPerson = `Hi I am {{.Name}}. {{.Age}} years old {{"\n"}}`

	// greetTemplate, err := template.New("greetingFromPerson").Parse(greetPerson)
	// if err != nil {
	// 	return
	// }
	greetTemplate := template.Must(template.New("greetingFromPerson").Parse(greetPerson))

	fmt.Println(greetTemplate.Name())
	greetTemplate.Execute(os.Stdout, filicity)
	greetTemplate.Execute(os.Stdout, oliver)

}
