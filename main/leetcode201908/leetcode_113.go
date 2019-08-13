package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	var per []int
	return getPathSum(root, 0, sum, per)
}

func getPathSum(root *TreeNode, now int, sum int, per []int) [][]int {
	var re [][]int
	if root.Left == nil && root.Right == nil {
		if now+root.Val == sum {
			item := append(per, root.Val)
			count := 0
			for e := range item {
				count += e
			}
			fmt.Println(count)
			fmt.Println(count == sum)
			re = append(re, item)
		}
		return re
	}
	var left, right [][]int
	if root.Left != nil {
		left = getPathSum(root.Left, now+root.Val, sum, append(per, root.Val))
	}
	if root.Right != nil {
		right = getPathSum(root.Right, now+root.Val, sum, append(per, root.Val))
	}
	if left != nil {
		re = append(re, left...)
	}
	if right != nil {
		re = append(re, right...)
	}
	return re
}

//func copyPath(per []int, appendValue int) []int {
//	item := make([]int, len(per))
//	copy(item, per)
//	item = append(item, appendValue)
//	return item
//}

func main() {
	node3 := &TreeNode{Val: 3}
	node2 := &TreeNode{Val: 2}
	node1 := &TreeNode{1, node2, node3}
	sum := pathSum(node1, 3)
	fmt.Println(sum)
}
