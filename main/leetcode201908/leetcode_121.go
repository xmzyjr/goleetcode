package main

func maxProfit1(prices []int) int {
	max := -1
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] > prices[i] {
				max = maxNumProfit1(max, prices[j]-prices[i])
			}
		}
	}
	return max
}

func maxNumProfit1(a, b int) int {
	if a > b {
		return a
	}
	return b
}
