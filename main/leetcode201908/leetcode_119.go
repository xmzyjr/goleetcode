package main

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return nil
	}
	result1 := []int{1}
	result2 := []int{1, 1}
	if rowIndex == 1 {
		return result1
	}
	if rowIndex == 2 {
		return result2
	}
	var result = result2
	for i := 1; i < rowIndex; i++ {
		temp := []int{1}
		for j := 1; j < len(result); j++ {
			temp = append(temp, result2[j]+result2[j-1])
		}
		temp = append(temp, 1)
		result = temp
	}
	return result
}
