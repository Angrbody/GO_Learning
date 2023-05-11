package singleton_test

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
}

var SingleInstance *Singleton
var once sync.Once

// Go的单例模式（懒汉），非常简单
func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create Obj!")
		SingleInstance = new(Singleton)
	})
	return SingleInstance
}

func TestSingleton(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("%x\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
