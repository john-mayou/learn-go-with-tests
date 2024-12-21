package sum

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	sums := []int{}
	for _, slice := range slices {
		sums = append(sums, Sum(slice))
	}
	return sums
}
