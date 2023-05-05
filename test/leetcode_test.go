package main

import (
	"fmt"
	"sort"
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
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
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

	sort.Ints(nums)
	for _, value := range nums {
		fmt.Println(value, " ")
	}

	sort.Slice(nums, func(i, j int) bool {
		tmpI, tmpJ := nums[i], nums[j]
		if tmpI < 0 {
			tmpI = -1 * tmpI
		}
		if tmpJ < 0 {
			tmpJ = -1 * tmpJ
		}
		return tmpI < tmpJ
	})
}

func FindGas(gas, cost []int) int {
	work := []int{}
	for i := range gas {
		work = append(work, gas[i]-cost[i])
	}

	for i, v := range work {
		if v < 0 {
			continue
		}

		cur := work[i]
		flag := true
		for j := (i + 1) % len(work); j != 0; j++ {
			cur += work[j]
			if cur < 0 {
				flag = false
				break
			}
		}
		if flag {
			return i
		}
	}
	return -1
}

func canCompleteCircuit(gas []int, cost []int) int {
	work := []int{}
	for i := range gas {
		work = append(work, gas[i]-cost[i])
	}

	curSum, totalSum, start := 0, 0, 0
	for i, v := range work {
		curSum += v
		totalSum += v
		if curSum < 0 {
			curSum = 0
			start = i + 1
		}
	}
	if totalSum < 0 {
		return -1
	}
	return start
}

func candy(ratings []int) int {
	res := len(ratings) // 每人至少一颗
	add := make([]int, len(ratings))
	// 从左到右遍历
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			add[i] = add[i-1] + 1
		}
	}

	// 从右到左遍历
	for i := len(ratings) - 2; i > 0; i-- {
		if ratings[i] > ratings[i+1] && add[i] <= add[i+1] {
			add[i] = add[i+1] + 1
		}
	}

	for _, v := range add {
		res += v
	}

	return res
}
