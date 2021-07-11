package main

import "sort"

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	var i int
	for i = 0; i < len(nums) && k > 0; i++ {
		if nums[i] < 0 {
			nums[i] = -nums[i]
			k--
		}
	}
	getSum := func(nums []int) int {
		var sum int
		for i := range nums {
			sum += nums[i]
		}
		return sum
	}
	if k > 0 && k%2 == 1 {
		sort.Ints(nums)
		nums[0] = -nums[0]
	}
	return getSum(nums)
}
