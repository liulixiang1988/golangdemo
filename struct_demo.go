package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func modify(per *person) {
	per.age = 10
}

func main() {
	per := &person{
		name: "lixiang",
		age:  26,
	}
	fmt.Println(per)
	modify(per)
	fmt.Println(per)

	a := &struct {
		name string
		age  int
	}{
		name: "longlong",
		age:  30,
	}

	fmt.Println(a)
}
