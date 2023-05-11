package utilalldone

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	// 输出调用此函数的协程id
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("the result is from %d", id)
}

func AllResponse() string {
	numOfRunner := 10
	// 使用buffered类型的channel避免协程资源浪费
	ch := make(chan string, numOfRunner) // 利用channel的非阻塞唤醒机制
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	// 等待所有结果都返回
	finalRet := ""
	for j := 0; j < numOfRunner; j++ {
		finalRet += <-ch + "\n"
	}
	return finalRet
}

func TestFirstResponse(t *testing.T) {
	t.Log("before:", runtime.NumGoroutine())
	t.Log(AllResponse())
	time.Sleep(time.Second * 1)
	t.Log("after:", runtime.NumGoroutine())
}
