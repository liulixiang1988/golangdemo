package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a = 65
	b := strconv.Itoa(a)
	fmt.Println(b)
	a, _ = strconv.Atoi(b) #返回多个值，注意
	fmt.Println(a)
}
