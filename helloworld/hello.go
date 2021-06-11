package main

import (
	"fmt"
	"log"
	"myFmt"

	"github.com/golang/example/stringutil"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(stringutil.Reverse("Hello World"))
	myFmt.Println("Hello world")
	log.Println("I am here")
}
