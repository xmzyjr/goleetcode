package exploretion02

func dominantIndex(nums []int) int {
	if nums == nil || len(nums) == 1 {
		return 0
	}
	max := func() int {
		max, index := nums[0], 0
		for k, v := range nums {
			if v > max {
				max = v
				index = k
			}
		}
		return index
	}()

	for k, v := range nums {
		if k != max {
			if v*2 > nums[max] {
				return -1
			}
		}
	}
	return max
}
