package main

import "fmt"

func main() {
	test := new(TreeNode)

	te(test)
}

func te(root *TreeNode) {
	fmt.Println(root.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return get(root.Left, root.Right)
}

func get(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if (left == nil || right == nil) || left.Val != right.Val {
		return false
	}
	return get(left.Left, right.Right) && get(left.Right, right.Left)
}
