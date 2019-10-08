package exploretion02

import "container/list"

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	stack := list.New()
	temp := root
	var re []int
	for stack.Len() != 0 || temp != nil {
		if nil != temp {
			stack.PushBack(temp)
			temp = temp.Left
		} else {
			back := stack.Back()
			node := back.Value.(*TreeNode)
			stack.Remove(back)
			re = append(re, node.Val)
			temp = node.Right
		}
	}
	return re
}
