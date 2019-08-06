package main

import (
	"fmt"
	"sort"
)

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	num := [][]*TreeNode{{root}}
	re := [][]int{{root.Val}}
	for i := 0; i < len(num); i++ {
		var temp []*TreeNode
		var numTemp []int
		for j := 0; j < len(num[i]); j++ {
			if num[i][j].Left != nil {
				temp = append(temp, num[i][j].Left)
				numTemp = append(numTemp, num[i][j].Left.Val)
			}
			if num[i][j].Right != nil {
				temp = append(temp, num[i][j].Right)
				numTemp = append(numTemp, num[i][j].Right.Val)
			}
		}
		if temp != nil {
			num = append(num, temp)
			re = append(re, numTemp)
		}
	}
	for i, j := 0, len(re)-1; i < j; i, j = i+1, j-1 {
		re[i], re[j] = re[j], re[i]
	}
	return re
}

func main() {
	s := []int{5, 2, 6, 3, 1, 4}
	reverse := sort.Reverse(sort.IntSlice(s))
	fmt.Println(reverse)

}
