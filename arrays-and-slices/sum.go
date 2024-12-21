package sum

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	sums := make([]int, len(slices))
	for i, nums := range slices {
		sums[i] = Sum(nums)
	}
	return sums
}
