package leetcode201909

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return get(root, root.Val)
}

func get(root *TreeNode, num int) int {
	if root.Left == nil && root.Right == nil {
		return num
	}
	left, right := 0, 0
	if root.Left != nil {

		left = get(root.Left, num*10+root.Left.Val)

	}
	if root.Right != nil {
		right = get(root.Right, num*10+root.Right.Val)
	}
	return left + right
}
