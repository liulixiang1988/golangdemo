package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	go func() {
		c1 <- "hello"
		fmt.Println("write \"hello\" done!")

		c1 <- "world"
		fmt.Println("write \"world\" done!")

		fmt.Println("write go to sleep")
		time.Sleep(3 * time.Second)
		c1 <- "channel"
		fmt.Println("write \"channel\" done!")
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Read wake up")

	msg := <-c1
	fmt.Println("Reader ", msg)

	msg = <-c1
	fmt.Println("Reader ", msg)

	msg = <-c1
	fmt.Println("Reader ", msg)
}
