package equaltest_test

import "testing"

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestEqual(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	// c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	// t.Log(a == c) //	长度不同会报错
	t.Log(a == d)
}

func TestBitClear(t *testing.T) {
	a := 7              // 0111 可读可写可执行
	a = a &^ Readable   // 读位置清零
	a = a &^ Writeable  // 写位置清零
	a = a &^ Executable // 可执行位置清零
	t.Log(a&Readable == Readable,
		a&Writeable == Writeable,
		a&Executable == Executable)
}
