package main

import "fmt"

type Node struct {
	value string
	nodes []Node
}

func main() {
	p := Node{value: "p"}
	g := Node{value: "g"}
	b := Node{value: "b", nodes: []Node{p, g}}
	q := Node{value: "q"}
	s := Node{value: "s"}
	k := Node{value: "k"}
	r := Node{value: "r", nodes: []Node{q}}
	a := Node{value: "a", nodes: []Node{r, s}}
	root := Node{value: "a", nodes: []Node{b, a, k}}

	fmt.Println(root)
	outline([]string{}, &root)

}

func outline(stack []string, n *Node) {
	stack = append(stack, n.value)
	if len(n.nodes) == 0 {
		fmt.Println(stack)
	}

	for _, v := range n.nodes {
		outline(stack, &v)
	}
}
