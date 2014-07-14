package main

import (
	"fmt"
)

func main() {
	if a := 1; a > 1 {
		fmt.Println("Hello")
	} else { //else必须在}后面
		fmt.Println(a)
	}

	s := "abc"
	for i, n := 0, len(s); i < n; i++ { // 常⻅的for 循环，⽀持初始化语句。
		fmt.Println(s[i])
	}
	n := len(s)
	for n > 0 { // 替代while (n > 0) {}
		println(s[n-1]) // 替代for (; n > 0;) {}
		n--
	}
	for false { // 替代while (true) {}
		println(s) // 替代for (;;) {}
	}
}
