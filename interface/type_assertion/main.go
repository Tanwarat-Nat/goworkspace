package main

func main() {
	//x.(T)
	//key.(string)
	// cd.([]byte)
	// val.(map)[string]string)
	// sr.(*Rounter)
	// rv.(map[string]string)
}

/*
x.(T) -> x คือตัว interface expression ก็คือ expression อะไรก็ได้ สุดท้ายค่าที่
return กลับมามันก็คือ interface หรือจะเป็น interface อย่างเดียวก็ได้

ส่วนข้างใน (T) มันคือ type ที่เราจะเอา interface ตัวนี้ x มา check กับข้างใน ()
ก็คือ check dynamic type ของ x (operan ตัวนี้) กับตัว type ที่ past เข้ามา
หรือที่จะเรียกว่า assert type


*/
