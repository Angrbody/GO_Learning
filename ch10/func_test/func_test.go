package functest

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

// 对某一种函数，计算函数操作的时长
// 输入是函数类型，返回值也是函数类型(嵌套函数)
// 类似于装饰者模型
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent: ", time.Since(start).Seconds())
		return ret
	}
}

// 一个运行比较慢的函数
func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

// 函数可以有多个返回值
func TestFunc(t *testing.T) {
	a, _ := returnMultiValues()
	// 下划线可以忽略其中的一个返回值
	t.Log(a)

	tsSF := timeSpent(slowFunc)
	t.Log(tsSF(10))
}
