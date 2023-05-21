package error_test

import (
	"errors"
	"testing"
)

// 定义预制的错误
var ErrorLessThanTwo error = errors.New("n should not be less than 2")
var ErrorLargerThanHundred error = errors.New("n should not be larger than 100")

func GetFibonacci(n int) ([]int, error) {
	// go强调快速失败 -> 检测放到程序前面
	if n < 2 {
		return nil, ErrorLessThanTwo
	}

	if n > 100 {
		return nil, ErrorLargerThanHundred
	}

	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

// 接口类
type error interface {
}

func TestGetFibonacci(t *testing.T) {
	// 采用错误检查机制
	if v, err := GetFibonacci(10); err != nil {
		if err == ErrorLessThanTwo {
			t.Error("need a larger number")
		}
		if err == ErrorLargerThanHundred {
			t.Error("need a smaller number")
		}
	} else {
		t.Log(v)
	}
}
