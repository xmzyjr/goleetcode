package main

import (
	"fmt"
)

/*

	给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
	在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。
	找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
*/
func maxArea(height []int) int {

	length := len(height)
	height1 := make([][]int, length, length)
	for k, v := range height1 {
		for k1 := range v {
			if k == k1 {
				height1[k][k1] = 0
			}
		}
	}

	return 0
}

// index1 = 0, index2 = len(height)-1
func getMaxArea(height []int, index1 int, index2 int) int {
	if index1 == index2 {
		return 0
	}
	c := index2 - index1
	k := min(height[index1], height[index2])

	area := c * k
	left := getMaxArea(height, index1+1, index2)
	right := getMaxArea(height, index1, index2-1)

	return max(area, max(left, right))
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	//height := make([][]int, 5,5)

	//for k, v := range height {
	//	for k1, v1 := range v {
	//		fmt.Println(k,v,k1,v1)
	//	}
	//}

	fmt.Println()

}
