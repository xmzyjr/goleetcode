package main

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	result1 := [][]int{{1}}
	result2 := [][]int{{1}, {1, 1}}
	if numRows == 1 {
		return result1
	}
	if numRows == 2 {
		return result2
	}
	for i := 1; i < numRows-1; i++ {
		temp := []int{1}
		for j := 1; j < len(result2[i]); j++ {
			temp = append(temp, result2[i][j]+result2[i][j-1])
		}
		temp = append(temp, 1)
		result2 = append(result2, temp)
	}
	return result2
}
