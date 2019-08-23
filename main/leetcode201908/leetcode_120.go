package main

import "strconv"

func minimumTotal(triangle [][]int) int {
	return getMiniMumTotal(triangle, make(map[string]int), 0, 0)
}

func getMiniMumTotal(num [][]int, me map[string]int, row, col int) int {
	if row == len(num)-1 {
		me[itoa(row, col)] = num[row][col]
		return num[row][col]
	}
	if me[itoa(row, col)] == 0 {
		minNum := min(getMiniMumTotal(num, me, row+1, col), getMiniMumTotal(num, me, row+1, col+1)) + num[row][col]
		me[itoa(row, col)] = minNum
		return minNum
	} else {
		return me[itoa(row, col)]
	}
}

func itoa(row, col int) string {
	return strconv.Itoa(row) + "|" + strconv.Itoa(col)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
