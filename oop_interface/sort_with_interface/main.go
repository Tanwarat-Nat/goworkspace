package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name string
	age  int
}

type byName []*Person // เอา type นี้มา satify interface

//type Interface interface {
// Len is the number of elements in the collection.
func (p byName) Len() int {
	return len(p)
}

// Less reports whether the element with index i
// must sort before the element with index j.
func (p byName) Less(i, j int) bool {
	return p[i].name < p[j].name
}

// Swap swaps the elements with indexes i and j.
func (p byName) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	p := []*Person{
		{"A", 22},
		{"B", 21},
		{"B", 24},
		{"C", 20},
		{"C", 20},
		{"A", 21},
		{"A", 22},
		{"B", 22},
	}
	printPerson(p)
	sort.Sort(byName(p))
	printPerson(p)
}

func printPerson(p []*Person) {
	fmt.Println("------------")
	for _, v := range p {
		fmt.Println((*v))
	}
}

/*
การ sort ในภาษาโก ถึงแม้ว่าการทำงานจริง สมมติว่าเราเขียนเว็ป เราก็จะไม่ sort ในตัวของ
backend สมมติว่าต้องการเก็บข้อมูลใน database แล้วเราดึงข้อมูลออกมา การดึงข้อมูลอออกมา
เราก็จะ sort by name, sort by date of register เราก็จะไป sort จาก query มาเลย
ไม่ต้องมาเสียเวลา เสีย resource ในการเขียน sort ในโปรแกรมของเรา

ถ้าเราต้องการส่งข้อมูล un-sort ไปให้ช้างหน้า แล้วข้างหน้าเขาจะเอาไป sort เอง อันนั้นก็
แล้วแต่ว่าทาง UI team เขาจะจัด state ในส่วนของ front end ยังไง ก็จะไม่ค่อยมีในโลก
ของการทำงานจริง ไม่ค่อยได้เขียน sort ในตัวของโปรแกรม

เราจะมาดูว่าถ้าเราจะทำ sort ในภาษาโกเราจะทำยังไงได้บ้าง นอกจากวิธีแรก ก็คือ
1.เขียน sort algorithm ของเราเองเลย
2.ไปใช้ที่เขาเตรียมมาให้เเล้วใน standard package
*/
