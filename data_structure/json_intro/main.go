package main

func main() {

}

// สัมพันธ์ใกล้ชิดกับ struct พอสมควร
// https://jsonplaceholder.typicode.com/
/*
ถ้าเราจะเก็บ json แบบนี้ เราจะไปเก็บในรูปแบบไหน
จะเป็น map รึเปล่า เป็น array ของ map มั้ย
อันนี้ก็เป็น map ตัวนึง, อันนี้ก็เป็น map ตัวนึง
แล้วก็เอามาเก็บใน array แต่มันเก็บได้แบบเดียวทั้งหมด
ฉะนั้น data structure ที่เหมาะสมจะทำงานกับ json คือ struct
เพระามันเก็บได้ทั้ง string int bool

Marshal = จัดระเบียบ คือ

101010101(คอมอ่าน) ---> text/csv/json/xml(คนออก)

ใน memory การเก็บข้อมูลจะเป็น 1010101 ถ้าเราต้องการ marshel
ก็คือจัดระเบียบมัน ให้กลายเป็นไฟล์อะไรที่คนสามารถอ่านได้
เป็นอะไรที่เราสามารถ transfer ไปทาง internet ได้
File/string/stream ก็คือจัดมันมาเรียงเป็นระเบียบกัน

Unmarshel = Unจัดระเบียบ

text/csv/json/xml(คนออก) ----> struct (101010101(คอมอ่าน))

แล้วถ้าที่เขาส่งมาเป็นในรูปของ Marshel แล้วเราต้องการกลับมัน
ให้เป็น data structure แบบของเรา คือ struct เก็บไว้ใน memory



*/
