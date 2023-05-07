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
	points := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
	}
	// 先按start大小排序
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})

	// 计算重叠数量
	res := 1
	win := points[0] // 初始窗口

	for i := 1; i < len(points); i++ {
		// 若不重叠
		if points[i][0] > points[i-1][1] {
			win = points[i]
			res++
			continue
		}
		// 若重叠，更新窗口
		win[0] = points[i][0]
	}

	fmt.Print(res)
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

// func canCompleteCircuit(gas []int, cost []int) int {
// 	work := []int{}
// 	for i := range gas {
// 		work = append(work, gas[i]-cost[i])
// 	}

// 	curSum, totalSum, start := 0, 0, 0
// 	for i, v := range work {
// 		curSum += v
// 		totalSum += v
// 		if curSum < 0 {
// 			curSum = 0
// 			start = i + 1
// 		}
// 	}
// 	if totalSum < 0 {
// 		return -1
// 	}
// 	return start
// }

// func candy(ratings []int) int {
// 	res := len(ratings) // 每人至少一颗
// 	add := make([]int, len(ratings))
// 	// 从左到右遍历
// 	for i := 1; i < len(ratings); i++ {
// 		if ratings[i] > ratings[i-1] {
// 			add[i] = add[i-1] + 1
// 		}
// 	}

// 	// 从右到左遍历
// 	for i := len(ratings) - 2; i > 0; i-- {
// 		if ratings[i] > ratings[i+1] && add[i] <= add[i+1] {
// 			add[i] = add[i+1] + 1
// 		}
// 	}

// 	for _, v := range add {
// 		res += v
// 	}

// 	return res
// }
