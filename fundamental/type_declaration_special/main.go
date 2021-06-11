package main

import "fmt"

type KG float64
type LB float64

func (lb LB) toKg() KG {
	// 2.2046226218 lb. = 1 kg.
	// 1 lb. = (1 kg * 1 lb) / 2.2046226218 lb
	return KG(lb / 2.2046226218)
}

func (kg KG) toLB() LB {
	// 1 kg = 2.2046226218 lb
	// 3 kg = 3 kg * 2.2046226218 lb
	return LB(kg * 2.2046226218)
}

func (kString KG) toString() string {
	return "kg"
}

func main() {
	// 1 kg = 2.20462262185 lb
	// 1 lb = 0.45359237 kg
	k := KG(3)
	b := LB(2.2046226218)

	fmt.Println(k == b.toKg())
	fmt.Println(k.toLB() > b)
	fmt.Printf("%v %v\n", k, k.toString()) // 1 kg.
	// fmt.Println(b.toString ()) => 2.2 lbs.

}
