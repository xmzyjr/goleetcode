package main

func isBalanced(root *TreeNode) bool {
	memory := make(map[*TreeNode]int, 50)
	return get(root, memory)
}

func get(root *TreeNode, m map[*TreeNode]int) bool {
	if root == nil {
		return true
	}
	left, right := true, true
	if root.Left == nil && root.Right == nil {
		m[root] = 1
	} else {
		if root.Left != nil {
			left = get(root.Left, m)
		}
		if root.Right != nil {
			right = get(root.Right, m)
		}
		if abs1(m[root.Left]-m[root.Right]) > 1 {
			return false
		}
		m[root] = max1(m[root.Left], m[root.Right]) + 1
	}
	return left && right
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs1(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
