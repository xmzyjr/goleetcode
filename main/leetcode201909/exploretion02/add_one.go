package exploretion02

func plusOne(digits []int) []int {
	if digits == nil {
		return nil
	}
	extra := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+extra == 10 {
			extra = 1
			digits[i] = 0
		} else {
			extra = 0
			digits[i]++
			break
		}
	}
	if extra == 1 {

		return append([]int{1}, digits...)
	} else {
		return digits
	}
}
