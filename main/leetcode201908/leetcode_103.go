package main

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	num := [][]*TreeNode{{root}}
	re := [][]int{{root.Val}}
	for i := 0; i < len(num); i++ {
		var temp []*TreeNode
		for j := 0; j < len(num[i]); j++ {
			if num[i][j].Left != nil {
				temp = append(temp, num[i][j].Left)
			}
			if num[i][j].Right != nil {
				temp = append(temp, num[i][j].Right)
			}
		}
		if temp != nil {
			num = append(num, temp)
		}
	}
	var flag bool = false
	for i := 0; i < len(num); i++ {
		var temp []int
		if !flag {
			for j := len(num[i]) - 1; j >= 0; j-- {
				temp = append(temp, num[i][j].Val)
			}
			flag = !flag
		} else {
			for j := 0; j < len(num[i]); j++ {
				temp = append(temp, num[i][j].Val)
			}
			flag = !flag
		}
		if temp != nil {
			re = append(re, temp)
		}
	}
	return re
}
