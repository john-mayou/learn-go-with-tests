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

func SumAllTails(slices ...[]int) []int {
	sums := make([]int, 0, len(slices))
	for _, nums := range slices {
		if len(nums) > 1 {
			sums = append(sums, Sum(nums[1:]))
		} else {
			sums = append(sums, 0)
		}
	}
	return sums
}
