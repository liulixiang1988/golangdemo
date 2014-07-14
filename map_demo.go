package main

import (
	"fmt"
)

func main() {
	a := map[string]struct {
		name string
		age  int
	}{
		"lixiang":  {"理想", 26},
		"longlong": {"龙珑", 26},
	}
	fmt.Println(a)
	var m map[string]string = make(map[string]string)
	m["lixiang"] = "lixiang"
	fmt.Println(m)
	delete(m, "lixiang")
	fmt.Println(m)
}
