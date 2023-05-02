package constant_test

import "testing"

// 连续常量的初始化
const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

// 连续位常量的初始化
const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConstantTry1(t *testing.T) {
	a := 7
	t.Log(a&Readable == Readable,
		a&Writeable == Writeable,
		a&Executable == Executable)
}
