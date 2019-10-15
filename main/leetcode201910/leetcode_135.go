package leetcode201910

func candy(ratings []int) int {
	if ratings == nil || len(ratings) == 0 {
		return 0
	}
	if len(ratings) == 1 {
		return 1
	}
	num := make([]int, len(ratings))
	num[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i-1] >= ratings[i] {
			num[i] = 1
		} else {
			num[i] = num[i-1] + 1
		}
	}
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i] < ratings[i-1] {
			num[i-1] = max(num[i-1], num[i]+1)
		}
	}
	return sum(num)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sum(num []int) int {
	sum := 0
	for _, v := range num {
		sum += v
	}
	return sum
}
