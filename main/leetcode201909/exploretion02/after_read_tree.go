package exploretion02

import "container/list"

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	var result []int
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() != 0 {
		back := stack.Back()
		node := back.Value.(*TreeNode)
		stack.Remove(back)
		result = append([]int{node.Val}, result...)
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
	}
	return result
}
