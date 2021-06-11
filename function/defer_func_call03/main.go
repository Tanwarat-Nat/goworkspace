package main

import (
	"fmt"
	"time"
)

func stopWatch(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("function %s: tooks time %v\n", name, time.Now().Sub(start).Seconds())
	}
}

func loadingImage() {
	defer stopWatch("loadingImage")()
	time.Sleep(2500 * time.Millisecond)
	fmt.Println("loadingImage: done")
}

func main() {
	loadingImage()
}
