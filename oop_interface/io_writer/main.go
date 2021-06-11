package main

import (
	"fmt"
)

type Notebook struct {
	content []byte
}

func (nb *Notebook) String() string {
	return string(nb.content)
}

func (nb *Notebook) Write(p []byte) (n int, err error) {
	nb.content = append(nb.content, p...)
	return len(p), nil
}

func main() {
	nb := Notebook{} // ใช้แบบนี้ได้เลย เพระา notebook เร่ implement interface ไว้
	fmt.Println(nb)
	fmt.Fprintln(&nb, "helloworld")
	fmt.Println("&nb", &nb)
}
