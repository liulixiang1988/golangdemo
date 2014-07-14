package main

import (
	"fmt"
)

const (
	A = iota
	B
	C = "C"
	D
	E = iota
	F
)

const (
	f = 1
	e = 3
	c = 2
	d = 3
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := a[:3]
	b[0] = 10
	fmt.Println("a[0]:", a)
	fmt.Println("E:", E)
	fmt.Println(c)
}
