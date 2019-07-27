package concurrency

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(data chan<- int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			data <- i
		}
		close(data)
		wg.Done()
	}()
}

func dataReceiver(data <-chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if value, ok := <-data; ok {
				fmt.Println(value)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	var data = make(chan int)

	wg.Add(1)
	dataProducer(data, &wg)

	wg.Add(1)
	dataReceiver(data, &wg)

	wg.Add(1)
	dataReceiver(data, &wg)

	wg.Wait()
}
