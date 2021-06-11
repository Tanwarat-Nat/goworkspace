package main

import (
	"encoding/json"
	"fmt"
)

type Todo struct {
	UserId    int
	Id        int
	Title     string
	Completed bool
}

func main() {
	/*
	   ชุดข้อมูลที่รับเข้ามาเป็น json การประกาศ ก้็ต้องจัดให้อยู่ในรูปแบบของ
	   json format ที่ถูกต้อง คือ ต้องใส่ [] ครอบ {} --> `[{ }]`
	   เมื่อ unmarshel แล้วจะเป็นในรููปแบบ array ของ Todo (array ครอบ Todo)
	*/
	var data = `[{ 
		"userId": 1,
		"id": 1,
		"title": "delectus aut autem",
		"completed": false
	},
	{
		"userId": 1,
		"id": 2,
		"title": "quis ut nam facilis et officia qui",
		"completed": false
	}]`
	// แต่ตรงนี้เราประกาศเป็น array ของ Todo แต่เราใช้ slice เพระายังไม่รู้ size
	dataStruct := []Todo{}
	v := &dataStruct
	fmt.Println(dataStruct)
	//
	json.Unmarshal([]byte(data), v)
	fmt.Println(dataStruct)
	fmt.Println(len(dataStruct))

	/* เราอ่านข้อมูลมาจากค่าๆนึง (หรืออ่านจาก website ก็ได้ คลิปถัดไปจะอ่านจาก web service)
	เราได้ข้อมูลมาชุดนึง เรา paste ข้อมูลมาชุดนี้ เก้บไว้ที่ struct ที่เราตั้งชื่อมันว่า Todo จากนั้น

	*/

}

/*Unmarshel = Unจัดระเบียบ

text/csv/json/xml(คนออก) ----> struct (101010101(คอมอ่าน))

ส่งข้อมูลแบบ marshel (text/csv/json/xml(คนอ่าน)) มา  เราต้องทำการ unmarshrl
ให้มันกลับมาอยู่ในรูปแบบ data structure type ที่เราต้องการ (struct 101010110 (คอมอ่าน))

*/
