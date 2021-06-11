package main

import "fmt"

func main() {
	// map literal style
	items := map[string]int{
		"pen":    4,
		"pencil": 5,
	}
	fmt.Println(items)
	fmt.Println(items["pencil"])

	// make function style
	orders := make(map[string]int)
	orders["pen"] = 16
	orders["pencil"] = 15
	fmt.Println(orders)
	fmt.Println(orders["pen"])

	// map is a reference to a hash table
	// data structuer copy by value
	// but reference always point to same hash table
	x := items
	x["ruler"] = 7
	x = orders
	fmt.Println("items =", items)
	fmt.Println("x =", x)
	delete(items, "pencil")

	fmt.Println("items =", items)
	items["xxx"]++            // += 1
	fmt.Println(items["xxx"]) // zero value

	animals := map[string]int{
		"dog": 9,
		"cat": 0,
	}
	v, ok := animals["cat"]
	if !ok {
		fmt.Println("not exit", v)
	} else {
		fmt.Println("exit", v)
	}

	tempt := map[string]int{
		"one":   1,
		"two":   0,
		"three": 0,
		"four":  0,
		"five":  0,
	}
	fmt.Println(len(tempt))
	for k, v := range tempt {
		fmt.Println(k, v)
	} // คำตอบจะไม่เรียงกันในแต่ละการ run
}

/*
Map = key value base
เป็น dat structure ที่เก็บ key และก็จะมี value แล้วถ้าต้องการข้อมูล
value ออกมา ก็ถามว่า ถ้า key อันนี้ value จะเป็นแบบใด

map คือ reference to hash table
hash table คือ data structure ที่สามารถ map key ไปเป็น value ได้
การทำงานคือ จะมีฟังก์ชั่นที่ชื่อว่า hash function ที่จะทำการ
compute index แล้วก็เอา value ตรงนั้นไปใส่ในพื้นที่ๆ index ที่เราได้

สมมติ มีพื้นที่เก๋็บข้อมูลอันนึงให้เป็น array มี index 0-n
ใน hash table key จะไม่ใช่ตัวเลข(0-n)
จะเป็น int, string, list, slice ก็ได้ ซึ่งจะต้องเปลี่ยนพวกนี้ให้มาเป็น key ก่อน
key จะต้องเป็นอะไรที่สามารถใช้เครื่องหมาย == ได้ในภาษาโก
เวลาใส่ key ตัวใหม่เข้ามามันจะมีการเช็คก่อนว่า key ชนกันมั้ย
key เลยต้อง compare ได้ พอเราพาส key เข้ามาแล้ว
hash function ก็จะรับ key เข้ามา สุดท้ายก็จะบอกเราว่า
เราจะเก็บข้อมูลตรงนั้นไว้ส่วนไหนของ buckets(ถัง)

ใช้ float เป็น key ไม่ดี ไม้ชัดเจน compare กันไม่ได้

การประกาศแบบ map literal style

name := map[type]type{
	key : value,
	key2 : value2,
}

การประกาศแบบ make function style

name := mame(map[type]type)
map[key] = value
map[key2] = value2

เวลาเรียกดู เรียกตาม key ---> fmt.Println(name[key])
เวลาลบ ใช้ ---> delete(name, "key")

*/
