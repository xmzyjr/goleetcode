package main

func flatten(root *TreeNode) {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return
	}
	getFlatten(root)
}

func getFlatten(root *TreeNode) (*TreeNode, *TreeNode) {
	if root.Left == nil && root.Right == nil {
		return root, root
	}
	var left, bottom, temp *TreeNode
	if root.Left != nil {
		left, bottom = getFlatten(root.Left)
	}
	temp = root.Right
	if left != nil {
		root.Left = nil
		root.Right = left
		bottom.Right = temp
	}
	if root.Right != nil {
		_, bottom = getFlatten(root.Right)
	}
	return root, bottom
}
