package exploretion02

import "sort"

func arrayPairSum(nums []int) int {
	if nums == nil {
		return 0
	}
	if len(nums) == 2 {
		return pairMin(nums[0], nums[1])
	}
	sort.Ints(nums)
	var sum int
	for i := 1; i < len(nums); i += 2 {
		sum += pairMin(nums[i], nums[i-1])
	}
	return sum
}

func pairMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}
