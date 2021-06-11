package main

import "fmt"

func main() {

	//var name type = expression
	var x int = 1 + 2
	var y string = "hello"
	var z float64 = 4.0

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)

	// var name type
	var i int
	var j float64
	var k string
	var l []string

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)

	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", j)
	fmt.Printf("%T\n", k)
	fmt.Printf("%T\n", l)
	fmt.Println(l == nil)

	// var multiple(name) type declaration
	var a, b, c int

	fmt.Printf("%#v, %T\n", a, a)
	fmt.Printf("%#v, %T\n", b, b)
	fmt.Printf("%#v, %T\n", c, c)

	// var multiple(name) = multiple(expression)
	var d, e, f = 1, 2.5, "hello"
	fmt.Printf("%#v, %T\n", d, d)
	fmt.Printf("%#v, %T\n", e, e)
	fmt.Printf("%#v, %T\n", f, f)

	// var name type
	// name = expression

	var m string
	m = "string"

	fmt.Println(m)
	fmt.Printf("%T\n", m)

}
