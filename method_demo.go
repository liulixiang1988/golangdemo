package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func (per *Person) Profile() {
	fmt.Println(per.name, per.age)
}

type INT int

func (i *INT) Increment() {
	*i += INT(100)
}

func main() {
	p := Person{name: "lixiang", age: 27}
	p.Profile()
	var a INT = 2
	a.Increment()
	fmt.Println(a)
}
