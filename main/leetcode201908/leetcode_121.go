package main

func maxProfit(prices []int) int {
	max := -1
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] > prices[i] {
				max = maxNumProfit(max, prices[j]-prices[i])
			}
		}
	}
	return max
}

func maxNumProfit(a, b int) int {
	if a > b {
		return a
	}
	return b
}
