package client

// 包的导入路径需要在go_path下，从src目录开始往下写
import (
	"ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacci(5))
	arr := []int{1, 2, 3, 4, 5}
	t.Log(series.GetSum(arr))
	t.Log(series.Square(5))
}
