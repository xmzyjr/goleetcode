package game_19_10_15

func balancedStringSplit(s string) int {
	if len(s) == 0 {
		return 0
	}
	sum := 0
	lCount := 0
	rCount := 0
	for _, v := range s {
		if v == 'R' {
			rCount++
			if rCount == lCount {
				sum++
				rCount, lCount = 0, 0
			}
		} else if v == 'L' {
			lCount++
			if rCount == lCount {
				sum++
				rCount, lCount = 0, 0
			}
		}
	}
	return sum
}
