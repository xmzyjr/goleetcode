package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 1
	}
	nodes := []*TreeNode{root}
	var index int
	max := -1e5 - 1
	maxIndex := 1
	for len(nodes) != 0 {
		index++
		sum := 0
		var temp []*TreeNode
		for _, v := range nodes {
			sum += v.Val
			if v.Left != nil {
				temp = append(temp, v.Left)
			}
			if v.Right != nil {
				temp = append(temp, v.Right)
			}
		}
		if sum > int(max) {
			max = float64(sum)
			maxIndex = index
		}
		nodes = temp
	}
	return maxIndex
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
