package maptest

import "testing"

// map初始化
func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))

	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))
}

// map访问
func TestGetNotExistKey(t *testing.T) {
	// 1.访问不存在的key对应的value，编译器默认给出0
	m1 := map[int]int{}
	t.Log(m1[1])

	// 2.存在key，对应value本来就是0
	m1[2] = 0
	t.Log(m1[2])

	// so 怎么区分以上两种情况？
	// 通过返回值来判断
	m1[3] = 0
	if v, ok := m1[3]; ok {
		t.Logf("key 3's value is %d", v)
	} else {
		t.Logf("key 3 is not exist.")
	}
}

// 遍历map
func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}

// map的value可以是一个方法
