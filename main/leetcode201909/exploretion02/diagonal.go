package exploretion02

func findDiagonalOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 {
		return nil
	}
	var result []int
	xLen := len(matrix)
	yLen := len(matrix[0])
	if xLen == 1 && yLen == 1 {
		return matrix[0]
	} else if xLen == 1 {
		return matrix[0]
	} else if yLen == 1 {
		for i := 0; i < xLen; i++ {
			result = append(result, matrix[i][0])
		}
		return result
	}
	leftDown, rightUp := []int{0, 0}, []int{0, 0}
	result = append(result, matrix[0][0])
	leftDown[0], leftDown[1] = leftDownPlus(leftDown[0], leftDown[1], xLen)
	rightUp[0], rightUp[1] = rightUpPlus(rightUp[0], rightUp[1], yLen)
	flag := true
	for leftDown[0] != rightUp[0] && leftDown[1] != rightUp[1] {
		var temp []int
		if flag {
			temp = []int{rightUp[0], rightUp[1]}
			for temp[0] != leftDown[0] && temp[1] != leftDown[1] {
				result = append(result, matrix[temp[0]][temp[1]])
				temp[0]++
				temp[1]--
			}
			result = append(result, matrix[leftDown[0]][leftDown[1]])
		} else {
			temp = []int{leftDown[0], leftDown[1]}
			for temp[0] != rightUp[0] && temp[1] != rightUp[1] {
				result = append(result, matrix[temp[0]][temp[1]])
				temp[0]--
				temp[1]++
			}
			result = append(result, matrix[rightUp[0]][rightUp[1]])
		}
		leftDown[0], leftDown[1] = leftDownPlus(leftDown[0], leftDown[1], xLen)
		rightUp[0], rightUp[1] = rightUpPlus(rightUp[0], rightUp[1], yLen)
		flag = !flag
	}
	result = append(result, matrix[xLen-1][yLen-1])
	return result
}

func leftDownPlus(x, y, xLen int) (int, int) {
	if x < xLen-1 {
		x++
	} else {
		y++
	}
	return x, y
}

func rightUpPlus(x, y, yLen int) (int, int) {
	if y < yLen-1 {
		y++
	} else {
		x++
	}
	return x, y
}
