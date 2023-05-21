package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64
	b = int64(a) // 不支持隐式类型转换

	var c MyInt
	c = MyInt(b) // 即使是同类型别名也不可以

	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// aPtr = aPtr + 1 // 不支持指针自增等操作运算
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
	if s == "" {
		println("string s is nil")
	}
}

// 类型的预定义值
// math.MaxInt64

// go不支持指针运算
// string是值类型，其默认的初始化值为空字符串，而不是null
