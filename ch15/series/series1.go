package series

func GetSum(nums []int) int {
	var sum int = 0
	for _, value := range nums {
		sum += value
	}
	return sum
}
