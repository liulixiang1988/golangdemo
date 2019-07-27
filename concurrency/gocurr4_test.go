package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func isCancelled(cancelChan <-chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func cancel_1(cancelChan chan<- struct{}) {
	cancelChan <- struct{}{}
}

func cancel_2(cancelChan chan<- struct{}) {
	close(cancelChan)
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{})

	for i := 0; i < 5; i++ {
		go func(i int, cancelChan <-chan struct{}) {
			for {
				if isCancelled(cancelChan) {
					break
				}
				time.Sleep(5 * time.Millisecond)
			}
			fmt.Println(i, "Done")
		}(i, cancelChan)
	}

	cancel_1(cancelChan)
	time.Sleep(time.Second * 1)
}
