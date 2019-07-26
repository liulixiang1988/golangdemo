package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func AsyncTimeDelay(timeDelay time.Duration) chan string {
	var retCh = make(chan string)
	go func() {
		fmt.Println("start")
		time.Sleep(timeDelay * time.Millisecond)
		retCh <- "ok"
		fmt.Println("done")
	}()
	return retCh
}

// TestSelectAndTimeOut 多路选择与超时
func TestSelectAndTimeOut(t *testing.T) {
	retCh := AsyncTimeDelay(100)
	select {
	case v, r := <-retCh:
		fmt.Println("v:", v, " r:", r)
	case <-time.After(time.Millisecond * 50):
		fmt.Println("timeout")
	}
}
