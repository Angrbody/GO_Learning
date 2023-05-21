package channelclose_test

import (
	"fmt"
	"sync"
	"testing"
)

// 生产者-消费者模式
// 通过 buffer 共享内存来交换数据
func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i // 向管道中写数据
		}
		// 数据传输结束后关闭管道
		close(ch)
		// ch <- 11，panic：send on closed channel
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			// 从管道中读数据，并判断通道是否关闭
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)

	wg.Add(1)
	dataReceiver(ch, &wg)

	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
}
