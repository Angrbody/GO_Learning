package map_test

import "testing"

// go中的 map可以放置方法（函数）
func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

// go中没有内置的set，可以map[int]bool 实现
func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	if mySet[n] {
		t.Logf("%d is exist.", n)
	} else {
		t.Logf("%d is not exist.", n)
	}

	// 输出集合中独立元素的个数
	mySet[2] = true
	t.Log(len(mySet))

	// 根据key删除键值对
	delete(mySet, 1)
	n = 1
	if mySet[n] {
		t.Logf("%d is exist.", n)
	} else {
		t.Logf("%d is not exist.", n)
	}
}
