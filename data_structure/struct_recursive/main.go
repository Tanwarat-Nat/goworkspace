package main

import "fmt"

//ถ้าจะ recursive type ตัวมันเอง เราก็จะต้อง take address ของมัน

type BinaryTree struct {
	value int
	left  *BinaryTree
	right *BinaryTree
}

func main() {
	root := BinaryTree{value: 2} // หลักที่เคยบอก ถ้าใส่ filed name ไม่ต้องบอกหมดก็ได้
	left := BinaryTree{value: 1}
	right := BinaryTree{value: 3}

	root.left = &left
	root.right = &right

	fmt.Println(root) // value = 2, L-R = ค่า pointer

	showDF(&root)

}

func showDF(node *BinaryTree) { // ถ้ามันเป็น nil มันจะ skip ใน if ทั้งหมด
	if node != nil {

		showDF(node.left)
		fmt.Println(node.value)
		showDF(node.right)
	}
}
