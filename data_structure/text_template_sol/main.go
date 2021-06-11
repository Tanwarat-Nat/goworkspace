package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

func upperCase(s string) string {
	return strings.ToUpper(s)
}

func main() {
	filicity := Person{"Filicity", 23}
	oliver := Person{"Oliver", 25}

	people := []Person{filicity, oliver}

	const greetPerson = `Hi I am {{.Name | upperCase}}. {{.Age}} years old {{"\n"}}`
	const greetPeople = `{{range .}}Hi I am {{.Name}}. {{.Age}} years old {{"\n"}}{{end}}`

	maps := make(template.FuncMap)
	maps["upperCase"] = upperCase

	greetTemplate := template.Must(template.New("greetingFromPerson").
		Funcs(maps).
		Parse(greetPerson))

	greetPeopleTemplate := template.Must(template.New("greetingFromPeople").Parse(greetPeople))

	fmt.Println(greetTemplate.Name())
	greetTemplate.Execute(os.Stdout, filicity)
	greetTemplate.Execute(os.Stdout, oliver)
	fmt.Println("===========================")
	greetPeopleTemplate.Execute(os.Stdout, people)

}
