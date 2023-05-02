package condition_test

import "testing"

func TestIfMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("1==1")
	}

	// if v,err := someFun(); err==nil {

	// } else {

	// }
}

// go 的switch分支支持多个条件判断
func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

// go switch不需要手动加break
func TestSwitchCaseCondition(t *testing.T) {
I:
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
			break I // 跳出指定的循环
		case i%2 == 1:
			t.Log("odd")
		default:
			t.Log("unknown")
		}
	}
}
