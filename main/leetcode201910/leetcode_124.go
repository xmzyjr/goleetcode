package leetcode201910

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum124 int

func maxPathSum(root *TreeNode) int {
	sum124 = 0
	maxRoot(root)
	return sum124
}

func maxRoot(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := max124(maxRoot(root.Left), 0)
	right := max124(maxRoot(root.Right), 0)
	newValue := root.Val + left + right

	sum124 = max124(sum124, newValue)
	return root.Val + max(left, right)
}

func max124(a, b int) int {
	if a > b {
		return a
	}
	return b
}
