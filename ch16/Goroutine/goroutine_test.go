package goroutine_test

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		// 协程调度的顺序和代码的逻辑顺序不一致
		// 对于每个协程来说，i为共享的变量
		// 这里的协程使用的是匿名函数(闭包)
		go func(i int) {
			fmt.Println(i)
		}(i) // 这里必须把i显示的传到协程中，且为值传递，地址不同，不会发生竞争
	}
	time.Sleep(time.Millisecond * 50)
}
