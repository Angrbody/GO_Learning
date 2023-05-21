package select_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func AsyncService() chan string {
	// 设置 buffer channel，否则channel会阻塞，直到有人来取得管道中的数据
	retCh := make(chan string, 1)
	// 新开一个协程，和主程序的协程并行执行
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret // 非buffer的channel，向channel中放数据
		fmt.Println("service exited")
	}()
	return retCh
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("task is done.")
}

func TestService(t *testing.T) {
	// 串行输出
	fmt.Println(service())
	otherTask()
}

// 异步获得数据
func TestAsynService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh) // 从channel中取数据
	time.Sleep(time.Second * 1)
}

// 使用select进行超时判断
// 否则会阻塞在等待管道读写上，影响性能
func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}
