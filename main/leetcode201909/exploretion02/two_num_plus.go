package exploretion02

func twoSum(numbers []int, target int) []int {
	if numbers == nil || len(numbers) == 0 {
		return nil
	}
	left, right := 0, len(numbers)-1
	for left < right && numbers[left]+numbers[right] != target {
		num := numbers[left] + numbers[right]
		if num > target {
			right--
		} else if num < target {
			left++
		} else {
			break
		}
	}
	if left >= right {
		return nil
	} else {
		return []int{left + 1, right + 1}
	}
}
