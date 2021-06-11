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

type customSort struct {
	Persons []*Person
	less    func(i, j *Person) bool
}

func (p customSort) Len() int {
	return len(p.Persons)
}

func (p customSort) Less(i, j int) bool {
	return p.less(p.Persons[i], p.Persons[j])
}

func (p customSort) Swap(i, j int) {
	p.Persons[i], p.Persons[j] = p.Persons[j], p.Persons[i]
}

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
	// sort.Sort(byName(p))
	// printPerson(p)

	// ถ้าจะไล่กลับ ก-ฮ -ฬ ฮ-ก ก็ใช้ sort.Reverse
	sort.Sort(sort.Reverse(customSort{Persons: p, less: func(i, j *Person) bool {
		if i.age != j.age { //ถ้าจะเอาอายุขึ้นก่อน ก็แค่ย้ายมันไปเป็นเงื่อนไขแรก
			return i.age < j.age
		}
		if i.name != j.name {
			return i.name < j.name
		}

		return false
	}}))
	printPerson(p)
}

func printPerson(p []*Person) {
	fmt.Println("------------")
	for _, v := range p {
		fmt.Println((*v))
	}
}

/*
define client code ก่อน ว่าเราจะไปเรียก code มันยังไง
*/
