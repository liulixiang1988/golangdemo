package main

import (
	"fmt"
)

func test(x [2]int) {
	fmt.Printf("x: %p\n", &x)
	x[1] = 100
}

func main() {
	var a = [2]int{}
	fmt.Printf("a: %p\n", &a)
	fmt.Println(a)
	test(a)
	fmt.Println(a)
}
