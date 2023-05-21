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

// 多参数输入
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5, 6))
}

func Clear() {
	fmt.Println("clear resources.")
}

// 延迟执行函数
func TestDefer(t *testing.T) {
	// 延迟执行的函数（函数退出前执行）
	// 多用于安全释放资源、锁等收尾工作
	defer Clear()

	t.Log("started")

	// 程序异常中断，表示不可修复的错误
	// 即便在这种情况下，defer修饰的函数也会执行
	panic("fatal error")

	// panic后面的普通代码是不会执行的
	// fmt.Println("end")
}
