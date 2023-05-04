package main

import (
	"fmt"
	"testing"
)

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

const (
	INT_MAX = int(^uint((0)) >> 1)
	INT_MIN = ^INT_MAX
)

func maxSubArray(nums []int) int {
	var res, cur int = nums[0], 0
	for _, value := range nums {
		// 如果加上当前的value还是正的，就延长子数组的长度
		if cur+value >= 0 {
			cur += value
			res = Max(cur, res)
		}

		// 如果加上当前的value变成负的了，就没有必要继续加
		if cur+value < 0 {
			res = Max(cur, res)
			cur = Max(value, 0)
		}
	}
	res = Max(res, cur)
	return res
}

func TestMaxSubArray(t *testing.T) {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	ret := maxSubArray(arr)
	t.Log(ret)
	fmt.Println(ret)
}

func Swap(x, y int) (int, int) {
	return y, x
}

func main() {
	nums := [9]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	var res, cur int = 0, 0
	for _, value := range nums {
		// 如果加上当前的value还是正的，就延长子数组的长度
		if cur+value >= 0 {
			cur += value
			res = Max(cur, res)
		}

		// 如果加上当前的value变成负的了，就没有必要继续加
		if cur+value < 0 {
			res = Max(cur, res)
			cur = Max(cur, 0)
		}
	}
	res = Max(res, cur)
	fmt.Println(res)

	var a, b int = 1, 2
	fmt.Printf("swap前的a和b：a = %d, b = %d", a, b)
	a, b = Swap(a, b)
	fmt.Printf("swap后的a和b：a = %d, b = %d", a, b)
}
