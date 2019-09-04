package main

func fib(N int) int {
	var num [31]int
	num[0] = 0
	num[1] = 1
	if N <= 1 {
		return num[N]
	}
	for i := 2; i <= N; i++ {
		num[i] = num[i-1] + num[i-2]
	}
	return num[N]
}
