package main

import "fmt"

const value int = 1e9 + 7

func numPrimeArrangements(n int) int {
	if n == 1 {
		return 1
	}
	num := getNum()
	fmt.Println(num)
	var count1, count2 int
	for i := 1; i <= n; i++ {
		if num[i] == 0 {
			count1++
		} else {
			count2++
		}
	}
	return func() int {
		if count1 > count2 {
			count1, count2 = count2, count1
		}

		re1, re2 := 1, 1
		for i := count1; i >= 1; i-- {
			re1 = (re1 * i) % value
		}

		for i := count2; i > count1; i-- {
			re2 = (re2 * i) % value
		}

		return ((re1 * re1) % value * re2) % value
	}()
}

func getNum() [101]int {
	var num [101]int
	num[0] = -1
	num[1] = -1
	for i := 2; i < len(num); i++ {
		if num[i] == 0 {
			for j := 2 * i; j < len(num); j += i {
				num[j] = -1
			}
		}
	}
	return num
}

func main() {
	fmt.Println(numPrimeArrangements(3))
}
