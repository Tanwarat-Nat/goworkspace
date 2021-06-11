package main

import (
	"encoding/json"
	"fmt"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Completed *bool  `json:"completed,omitempty"`
}

func main() {
	var data = `[{ 
		"userId": 1,
		"title": "delectus aut autem",
		"completed": false
	},
	{
		"userId": 1,
		"id": 2,
		"title": "quis ut nam facilis et officia qui",
		"completed": true
	},
	{
		"userId": 2,
		"id": 5
	}
]`
	dataStruct := []Todo{} // แต่ตรงนี้เราประกาศเป็น slice เพระาไม่รู้ size
	v := &dataStruct
	fmt.Println(dataStruct)
	json.Unmarshal([]byte(data), v)
	fmt.Println(dataStruct)
	fmt.Println(len(dataStruct))
	// dataStruct[0].Completed = true
	// dataStruct[1].Completed = true

	// data structure --> std.Output
	result, err := json.MarshalIndent(dataStruct, "", "    ")
	if err != nil {
		return
	}
	fmt.Println(string(result))

}

/*
Marshal = จัดระเบียบ คือ

101010101(คอมอ่าน) ---> text/csv/json/xml(คนอ่าน)

ใน memory การเก็บข้อมูลจะเป็น 1010101 ถ้าเราต้องการ marshel
ก็คือจัดระเบียบมัน ให้กลายเป็นไฟล์อะไรที่คนสามารถอ่านได้
เป็นอะไรที่เราสามารถ transfer ไปทาง internet ได้
*/

/* เราอ่านข้อมูลมาจากค่าๆนึง (หรืออ่านจาก website ก็ได้ คลิปถัดไปจะอ่านจาก web service)
เราได้ข้อมูลมาชุดนึง เรา paste ข้อมูลมาชุดนี้ เก้บไว้ที่ struct ที่เราตั้งชื่อมันว่า Todo จากนั้น
เราก็ทำการ update Todo ว่ามัน completed แล้วนะ มัน true...
สิ่งที่เราจะทำต่อก็คือ ทำการเขียน data ที่ update ตรงนี้กลับไปเป็น json ใหม่
เขียนใส่ใน file ก็ได้ หรือเขียนใส่ standard output ก็ได้
*/
