package main

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	return getHasPathSum(root, 0, sum)
}

func getHasPathSum(root *TreeNode, now int, sum int) bool {
	if root.Left == nil && root.Right == nil {
		return root.Val+now == sum
	}
	left, right := false, false
	if root.Left != nil {
		left = getHasPathSum(root.Left, now+root.Val, sum)
	}
	if root.Right != nil {
		right = getHasPathSum(root.Right, now+root.Val, sum)
	}
	return left || right
}
