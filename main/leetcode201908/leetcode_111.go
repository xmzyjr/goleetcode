package main

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	queue := []*TreeNode{root}
	var count int
out:
	for len(queue) > 0 {
		var queue1 []*TreeNode
		for _, value := range queue {
			if value.Left == nil && value.Right == nil {
				break out
			}
			if value.Left != nil {
				queue1 = append(queue1, value.Left)
			}
			if value.Right != nil {
				queue1 = append(queue1, value.Right)
			}
		}
		count++
		queue = queue1
	}
	return count + 1
}
