package sharemem

import (
	"sync"
	"testing"
	"time"
)

// 线程不安全
func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

// 线程安全
func TestCounterThreadSafe(t *testing.T) {
	counter := 0
	var mut sync.Mutex
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

// 使用waitGroup，操作更方便灵活
func TestCounterWaitGroup(t *testing.T) {
	counter := 0
	var mut sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 5000; i++ {
		wg.Add(1) // 每有一个协程就需要 add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done() // 协程操作完成就 done()
		}()
	}
	// time.Sleep(1 * time.Second)
	wg.Wait() // 代替固定的sleep时间
	t.Logf("counter = %d", counter)
}
