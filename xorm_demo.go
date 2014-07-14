package main

import (
	"fmt"

	"db"
)

func main() {
	fmt.Println("插入数据：")
	err := db.NewDetail(2008, 17.04)
	if err != nil {
		fmt.Println("插入数据失败", err)
	}
}
