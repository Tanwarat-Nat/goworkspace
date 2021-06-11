package main

import (
	"errors"
	"fmt"
	"os"
)

type MyError struct {
	msg string
}

func (m MyError) Error() string {
	return m.msg
}

func main() {
	//ioutil.ReadFile()
	os.Open("/a/b/c/test.go")
	var err error
	err = errors.New("sdf")
	fmt.Printf("%T, %#v\n", err, err)
	fmt.Println(errors.New("sdf") == errors.New("sdf"))

	err = MyError{"This is an error"}
	fmt.Printf("%T, %#v\n", err, err)
}

/*
error ที่บอกว่าเป็น interface จริงๆแล้วมันเป็นอะไร

ในการทำงานจริงเราไม่ค่อย satify error interface สักเท่าไหร่

ส่วนใหญ่เราจะใช้ error.new Errorf bla bla bla

*/
