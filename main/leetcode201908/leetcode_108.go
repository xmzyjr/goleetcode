package main

func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	root := new(TreeNode)
	mid := len(nums) / 2
	root.Val = nums[mid]
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}
