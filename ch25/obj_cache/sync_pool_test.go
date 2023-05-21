package objcache

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create a new object.")
			return 100
		},
	}

	v := pool.Get().(int) // 使用断言
	fmt.Println(v)

	pool.Put(3)
	runtime.GC() // 手动调用GC

	time.Sleep(time.Second * 1)
	v1, _ := pool.Get().(int)
	fmt.Println(v1)
}

func TestSyncPoolInMultiGoroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create a new object.")
			return 10
		},
	}

	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Println(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// 合并区间
func merge(intervals [][]int) [][]int {

}
