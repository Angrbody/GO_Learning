package cancelbyclose

import (
	"fmt"
	"testing"
	"time"
)

func isCancelled(cancelChan chan struct{}) bool {
	// 检查是否从channel上收到数据
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{} // 发送一个空消息表示取消任务
}

func cancel_2(cancelChan chan struct{}) {
	close(cancelChan)
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{})
	// 5个协程，只有任务被取消的时候才会停止循环
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled.")
		}(i, cancelChan)
	}
	cancel_1(cancelChan) // 只能取消一个协程
	cancel_2(cancelChan) // 通过广播机制取消所有协程
	time.Sleep(time.Second * 1)
}
