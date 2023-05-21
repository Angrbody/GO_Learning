package testarray

import "testing"

func TestArrayInt(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 2, 3, 4}
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr3)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}

	// 类似 for_each 操作
	// 可以同时获得下标和值
	for idx, e := range arr3 {
		t.Log(idx, e)
	}
}

// 数组截取
// a[1:2] 左闭右开

func TestArraySection(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5}
	a_sec := a[3:4]
	t.Log(a_sec)
}

// 切片 slice ： 可变长数组
/*
struct {
	int *ptr
	int size
	int capacity
}
*/
