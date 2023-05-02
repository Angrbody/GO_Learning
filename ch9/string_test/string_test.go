package stringtest

import (
	"strconv"
	"testing"
)

func TestString(t *testing.T) {
	var s string
	t.Log(len(s))

	s = "hello"
	t.Log(len(s))

	s = "\xE4\xB8\xA5" // 可以存储任何二进制数据
	t.Log(len(s))

	t.Log(s)
	s = "中"
	t.Log(len(s))

	c := []rune(s)

	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

func TestStringLen(t *testing.T) {
	str := "hello" + "世界"
	t.Log(str)

	t.Log(len(str)) // 11 = 5 + 3 + 3，其中一个汉字要占3个字节

	t.Log(len([]rune(str))) // 7，rune数组表示 unicode字符，一个汉字占一位

	numStr := "1233445"
	val, err := strconv.Atoi(numStr)
	if err != nil {
		t.Log("strconv error!")
	} else {
		t.Logf("convert succeed!, val = %d", val)
	}
}
