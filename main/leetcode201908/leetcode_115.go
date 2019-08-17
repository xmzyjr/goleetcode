package main

func numDistinct(s string, t string) int {
	if s == t || len(t) == 0 {
		return 1
	}
	if s == "" || len(s) < len(t) {
		return 0
	}
	left, right := 0, 0
	if s[0] == t[0] {
		left = numDistinct(s[1:], t[1:])
	}
	right = numDistinct(s[1:], t)
	return left + right
}

func numDistinct1(s, t string, index1, index2, lens, lent int) int {
	if index2 == lent && index1 >= index2 {
		return 1
	}
	if (index1 == lens && index2 < lent) || lens-index1 < lent-index2 {
		return 0
	}
	left, right := 0, 0
	if s[index1] == t[index2] {
		left = numDistinct1(s, t, index1+1, index2+1, lens, lent)
	}
	right = numDistinct1(s, t, index1+1, index2, lens, lent)
	return left + right
}

func numDistinct3(s string, t string) int {
	if s == t || len(t) == 0 {
		return 1
	}
	if s == "" || len(s) < len(t) {
		return 0
	}
	lens, lent := len(s), len(t)
	num := make([][]int, lens+1)
	for k := range num {
		num[k] = make([]int, lent+1)
	}
	for i := 0; i < lens+1; i++ {
		if i >= lent {
			num[i][lent] = 1
		}
	}
	for i := lens - 1; i >= 0; i-- {
		for j := lent - 1; j >= 0; j-- {
			if lens-i < lent-j {
				num[i][j] = 0
			} else {
				if s[i] == t[j] {
					num[i][j] += num[i+1][j+1]
				}
				num[i][j] += num[i+1][j]
			}
		}
	}
	return num[0][0]
}
