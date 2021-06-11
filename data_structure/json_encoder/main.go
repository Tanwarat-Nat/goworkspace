package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     Geo    `json:"geo"`
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

type Users []struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
	Company  Company `json:"company"`
}

func main() {
	//https://jsonplaceholder.typicode.com/users
	//https://charnnarong.github.io/json-to-go-struct/
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		return
	}
	// json.Decode --> data structure
	jsonDecoder := json.NewDecoder(resp.Body)
	dataStruct := Users{}
	jsonDecoder.Decode(&dataStruct)
	resp.Body.Close()

	fmt.Println(len(dataStruct))
	dataStruct[0].Name = "Tanwarat"

	// json.Encode --> json
	jsonEncoder := json.NewEncoder(os.Stdout)
	jsonEncoder.Encode(dataStruct)

}

/*
encoder เอาอะไรอย่างนึงที่เราไม่รู้จัก code ไปในรูปของ json
struct --> json

decode เรารู้จักว่า json เป็นยังไง แต่คนอื่นไม่รู้ เราต้องทำให้เป็น struct
json --> struct
json.Decode --> data structure
มีไฟล์อันนึงจาก local หรือ respond ที่มาจาก web ก็ได้ เราก็จะมี file reader อันนึง
ซึ่ง file reader อันนี้มันเป็น abstract layer นึงที่เราจะ read จากไฟล์ก็ได้
จาก respond ก็ได้ หรือจะ read จาก que ก็ได้ พอเราได้ Reader มาแล้ว
json ก็จะ abstract ไปอีกอันนึง ก็คือ Json Decoder ไม่รู้จัก stream , file อะไรทั้งนั้น
รู้จักอย่างเดียวคือ file reader รู้จักแล้วก็สามารถเอาไฟล์พวกนี้มา decode ได้
ไฟล์ทั้งหมดนั้นก็ต้องอยู่ในรูป json แล้วทำการ decode ไปเป็น struct instance ที่เราต้องการ



*/
