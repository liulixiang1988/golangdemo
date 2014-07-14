package main

import (
	"fmt"
)

func A(a, b, c int) (r int, e string) {
	e = "hello"
	r = a + b + c
	return
}

func S(a ...int) (r int) {
	fmt.Println(r)
	for _, b := range a {
		r += b
	}
	return
}

func closure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func panic_demo() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover demo")
		}
	}()
	panic("panic demo")
}

func main() {
	f, s := A(1, 2, 3)
	fmt.Println(f, s)
	a := S(1, 2, 3, 4, 5)
	fmt.Println(a)

	b := closure(1)
	fmt.Println(b(2))

	defer fmt.Println(a)
	defer fmt.Println(b(2))

	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}

	panic_demo()
}
