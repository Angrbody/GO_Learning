package series

import "fmt"

func init() {
	fmt.Println("init1")
}

// 可以有两个同名init函数，并且依次进行
func init() {
	fmt.Println("init2")
}

func GetFibonacci(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}

// 小写开头的函数不能被外部访问
// func square(n int) int {}
func Square(n int) int {
	return n * n
}
