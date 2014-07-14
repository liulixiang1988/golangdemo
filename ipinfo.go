package main

import (
	"github.com/yinheli/qqwry"
	"log"
)

func main() {
	q := qqwry.NewQQwry("qqwry.dat")
	q.Find("180.89.94.90")
	log.Printf("ip:%v, 国家：%v, City:%v", q.IP, q.Country, q.City)
}
