package exploretion02

func findMaxConsecutiveOnes(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	max := 0
	temp := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			temp++
		} else {
			if temp != 0 {
				if temp > max {
					max = temp
				}
				temp = 0
			}
		}
	}
	if temp > max {
		return temp
	}
	return max
}
