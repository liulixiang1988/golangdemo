// concurrency go并发学习
// 使用go test xx.go
// 使用go test -v xx.go 查看详情
// 使用go test 执行所有测试
// 使用go test -v -run="TestA" 执行匹配的测试
package concurrency

import (
	"fmt"
	"sync"
	"testing"
)

// TestGoroutine goroutine与waitgroup
func TestGoroutine(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
	t.Log("testing")
}

// sync.Mutex用来做共享控制，实现计数器
func TestMutex(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup

	var counter = 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			defer mut.Unlock()
			mut.Lock()

			counter++
		}()
	}
	wg.Wait()
	t.Logf("counter: %d", counter)
}
