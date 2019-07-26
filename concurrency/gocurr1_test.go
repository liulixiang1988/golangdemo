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

// TestGoroutine 使用测试
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
