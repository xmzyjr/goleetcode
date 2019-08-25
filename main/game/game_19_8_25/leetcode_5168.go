package main

import (
	"fmt"
)

func numSmallerByFrequency(queries []string, words []string) []int {
	numQueries, numWords, ans := make([]int, len(queries)), make([]int, len(words)), make([]int, 0, len(queries))
	for i, v := range queries {
		numQueries[i] = f(v)
	}
	for i, v := range words {
		numWords[i] = f(v)
	}

	for i := 0; i < len(numQueries); i++ {
		count := 0
		for j := 0; j < len(numWords); j++ {
			if numWords[j] > numQueries[i] {
				count++
			}
		}
		ans = append(ans, count)
	}
	return ans
}

func f(str string) int {
	if len(str) <= 1 {
		return len(str)
	}
	var num [26]int
	for _, v := range str {
		num[v-'a']++
	}
	for e := range num {
		if num[e] != 0 {
			return num[e]
		}
	}
	return 0
}

func main() {
	fmt.Println(numSmallerByFrequency([]string{"bbb", "cc"}, []string{"a", "aa", "aaa", "aaaa"}))
}
