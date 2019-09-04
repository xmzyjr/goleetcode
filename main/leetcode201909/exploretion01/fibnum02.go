package main

func climbStairs(n int) int {
	var num [100]int
	num[1] = 1
	num[2] = 2
	if n <= 2 {
		return num[n]
	}
	for i := 3; i <= n; i++ {
		num[i] = num[i-1] + num[i-2]
	}
	return num[n]
}
