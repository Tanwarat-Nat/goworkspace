package main

import "fmt"

type Person struct {
	Name    string
	Surname string
}

func (p *Person) Fullname() string {
	return p.Name + " " + p.Surname
}

type Employee struct {
	Person
	Id     string
	Office string
}

func (e *Employee) Detial() string {
	return "ID :" + e.Id + ". Office :" + e.Office + ". Fullname :" + e.Fullname()
}

func (e *Employee) IsSameOffice(other *Employee) bool {
	return e.Office == other.Office
}

type Programmer struct {
	Employee
	Language []string
}

func (p *Programmer) Detial() string {
	return "Programmer :" + p.Employee.Detial()
}

type Tester struct {
	Employee
	Tools []string
}

func (t *Tester) Detial() string {
	return "Tester :" + t.Employee.Detial()
}

func main() {
	david := Person{
		Name:    "David",
		Surname: "Wright",
	}
	empDavid := Employee{
		Person: david,
		Id:     "123",
		Office: "Thailand",
	}
	progDavid := Programmer{
		Employee: empDavid,
		Language: []string{"Golang", "Java", "C++"},
	}
	fmt.Println("progDavid: ", progDavid)
	fmt.Printf("%+v\n", progDavid)

	oliver := Person{
		Name:    "Oliver",
		Surname: "Quenn",
	}
	empOliver := Employee{
		Person: oliver,
		Id:     "456",
		Office: "Thailand",
	}

	testerOliver := Tester{
		Employee: empOliver,
		Tools:    []string{"Robot"},
	}

	fmt.Println("testerOliver: ", testerOliver)
	fmt.Printf("%+v\n", testerOliver)

	fmt.Println("emp: sameoffice: ", empDavid.IsSameOffice(&empOliver))
	fmt.Println("prog: sameoffice: ", progDavid.IsSameOffice(&(testerOliver.Employee)))

	fmt.Println(progDavid.Fullname())
	fmt.Println(progDavid.Detial())

	daviaDetial := progDavid.Detial()
	fmt.Println("daviaDetial: ", daviaDetial)

	isSameOffice := (*Employee).IsSameOffice
	fmt.Println(isSameOffice(&empDavid, &empOliver))
	//fmt.Println((*Employee).IsSameOffice(&empDavid, &empOliver))

}

/*
ใน struct เราสามารถ extract method value ออกมาได้
แล้วเราจะเอา value ตรงนั้นพาสไปหา function อื่น หรือจะไปเรียกทีหลังก็ได้

อย่างเวลาเราเรียก david detial
*/
