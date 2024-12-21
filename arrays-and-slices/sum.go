package sum

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	sums := make([]int, 0, len(slices))
	for _, nums := range slices {
		sums = append(sums, Sum(nums))
	}
	return sums
}
