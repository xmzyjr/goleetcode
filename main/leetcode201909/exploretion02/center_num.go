package exploretion02

func pivotIndex(nums []int) int {
	if nums == nil || len(nums) <= 1 {
		return -1
	}
	sum := func() int {
		sum := 0
		for _, v := range nums {
			sum += v
		}
		return sum
	}()
	var left, right int
	for k, v := range nums {
		if k == 0 {
			left = 0
		} else {
			left += nums[k-1]
		}

		right = sum - v - left
		if right == left {
			return k
		}
	}
	return -1
}
