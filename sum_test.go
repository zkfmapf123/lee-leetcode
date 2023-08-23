package sum

import (
	"testing"
)

func twoSum(nums []int, target int) []int {
	numIndexMap := make(map[int]int) // Map to store number-index pairs

	for i, num := range nums {
		complement := target - num

		if index, found := numIndexMap[complement]; found {

			return []int{index, i} // Found a solution
		}

		numIndexMap[num] = i
	}

	return nil // No solution found
}

func TestTwoSum(t *testing.T) {
	twoSum([]int{2, 7, 11, 15}, 9)
}
