package main

import "fmt"

type Sword struct {
	name string
}

func (s Sword) detial() string {
	return "Super cold Iced sword"
}

func (s Sword) cost() int {
	return 100
}

type Bow struct {
	name string
}

func (b Bow) detial() string {
	return "Magic Fire bow"
}

func (b Bow) cost() int {
	return 50
}

type Item interface {
	cost() int
}

type Weapon interface {
	detial() string
	Item // embeded จะแบบนี้ก้ได้ --> cost() int
}

func attack(w Weapon) {
	fmt.Printf("attack by: %s. Item cost: %d\n", w.detial(), w.cost())
}

func main() {
	sword := Sword{name: "Icy-sword"}
	bow := Bow{name: "Fire-bow"}

	var w Weapon
	w = sword

	attack(w)
	w = bow
	attack(w)
}

/*เรามี interface อันนึงคือ weapon เราทำการ satisfy interface ตรงนี้
ถ้าเราใช้คำว่า implement ก็จะไม่ถูกซะทีเดียว เพระาว่าถ้าเป็นจาวา เวลาคุณมี interface
แล้วคุณเอา class ตรงนั้นไป implement interface

สำหรับในโกแล้ว เราจะใช้คำว่า implement ซะทีเดียวก็ไม่ได้นะ เพราะว่าจริงๆแล้วเนี่ย struct
ไม่ได้ทำการ implement interface อันนี้แล้วเท่านั้นเอง เพียงแค่ struct มันแค่มี
method ที่ชื่อว่า detial() ซึ่ง method ตรงนี้มันมี signature อันเดียวกันกับ
interface ตรงนี้ แล้วใน concept ของ type การที่ attack รับ parameter
ที่เป็น interface ของ weapon เข้ามา อันนี้คือ compile time นะ
compiler จะรู้ว่า sord  ของเรา struct ตรงนี้มันมี method อันเดียวกันกับที่
interface อันนี้บอกไว้ ดังนั้น attack ตรงนี้มันก้ยอมให้ object ที่เป็น sord, bow
ที่ทำการ satisfy interface สามารถพาร์ทเข้ามาในฟังก์ชั่นอันนี้ได้

การเท่ากันของ interface
ถ้าเรามี interface อันนึง แล้วเรามี object อยู่ 2 ตัว
ตัวแรก เป็น type struct อันนึง
อีกตัว ก็เป็น type struct อีกอันนึง
แต่ว่า struct ตัวนี้ไม่ได้ทำการ satisfy interface ซึ่งต้องการ 3  ตัว
มันก็ไม่สามารถ assign เข้ามาหา interface ตรงนี้ได้
ต่างกันกับตรงนี้ที่มีครบ มันจะสามารถ satisfy ได้ก็ต่อเมื่อ ทุกๆ method มันครบ
ทุกๆ method ใน struct มีอยู่ใน interface   ที่เขาประกาศไว้ อันนั้นถึงจะได้
ขาดไปตัวใดตัวนึงก็ไม่ได้

struct ต้อง satisfy interface ให้ครบทุกตัว interface ห้ามเหลือ
แต่ struct เหลือได้ จะถือว่าเป็น extra
struct	interface
 1 ----- 1
 2 ----- 2
 3 ----- 3 interface ต้องจับครบทุกตัว ห้ามเหลือ
 4 ----- struct เหลือได้ ถือเป็น extra

interface  embedded ได้เหมือนกัน คล้ายๆกับ struct
พอ embeded แล้วจะ satisfy swod, bow ไม่ได้
เพระาว่า weapon ของเรามันจะไปเท่ากับ swod ไม่ได้ เพราะมันมีตั้ง 2 method
คือ detial() กับ cost() แต่ว่า sword กับ bow เรามีแค่ detial() ไม่มี cost()
ดังนั้นเราก็ต้องใี cost()อันนึงให้มัน

- interface จะ satisfy ได้ก็ต่อเมื่อเรามี method ให้ครบ มันถึง satisfy ได้
- interface มันสามารถ embedded กันได้ มันก็คือการสร้าง interface ขึ้นมาใหม่
โดยการรวมกับ interface ที่มีอยู่แล้ว

จุดอันนึงที่น่าสังเกตก็คือว่า ในภาษาโกแทนที่เราประกาศ interface ขึ้นมาอันนึงบางที่
เราไม่จำเป็นต้องเข้าไปแก้ไข struct detial ของ exiting object ที่เรามีอยู่แล้ว
เป็น code ที่เรามีอยู่แล้ว ถ้าสมมติว่าตอนแรกไม่มี Item interface
แล้วใน weapon interface มี detial() กับ cost() พอวันนึงโปรแกรมมันโตขึ้นเรื่อย
เราเห็นร design ในโปรแกรมของเราว่าเห้ยย ทุกอย่างมันมี cost นะ
เรา abstract interface ตรงนี้ขึ้นมา โดยการสร้างขึ้นมาอีก interface นึง
Utem interface แล้ว cost() ตรงนี้ก็ให้มันเป็น Item ซะ
เราทำแค่สร้าง interface ใหม่ แล้วแก้ interface เก่า โดยที่ไม่ได้ไปแตะ sword กับ bow เลย
โปรแกรมของเราก้ยังทำงานถูกต้องเหมือนกัน

ถ้าเป็นภาษาอื่น พอเราแก้ไข interface ทีนึงก็ต้องไปเช็คดูว่า class ตรงนั้น
implement interface นู้นนี้นั่นรึเปล่า

*/
