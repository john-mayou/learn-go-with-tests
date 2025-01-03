package sum

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("returns sum of slice elements", func(t *testing.T) {
		assertEqualInt(t, Sum([]int{}), 0)
		assertEqualInt(t, Sum([]int{1, 2}), 3)
		assertEqualInt(t, Sum([]int{1, 2, 3, 4, 5}), 15)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("returns sums of slices", func(t *testing.T) {
		assertEqualIntSlice(t, SumAll(), []int{})
		assertEqualIntSlice(t, SumAll([]int{}), []int{0})
		assertEqualIntSlice(t, SumAll([]int{1, 2}, []int{3, 4, 5}), []int{3, 12})
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("returns sum of slices tails", func(t *testing.T) {
		assertEqualIntSlice(t, SumAllTails(), []int{})
		assertEqualIntSlice(t, SumAllTails([]int{}), []int{0})
		assertEqualIntSlice(t, SumAllTails([]int{1}), []int{0})
		assertEqualIntSlice(t, SumAllTails([]int{1, 2}), []int{2})
		assertEqualIntSlice(t, SumAllTails([]int{2, 3, 4}, []int{5, 6, 7}), []int{7, 13})
	})
}

func assertEqualInt(t testing.TB, actual, expected int) {
	t.Helper()
	if actual != expected {
		t.Errorf("expected %d but got %d", expected, actual)
	}
}

func assertEqualIntSlice(t testing.TB, actual, expected []int) {
	t.Helper()
	if !slices.Equal(actual, expected) {
		t.Errorf("expected %v but got %d", expected, actual)
	}
}
