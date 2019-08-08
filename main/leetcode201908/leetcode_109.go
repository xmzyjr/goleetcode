package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	left, mid, right := findMid(head)

	return &TreeNode{mid.Val, sortedListToBST(left), sortedListToBST(right)}
}

func findMid(root *ListNode) (left *ListNode, mid *ListNode, right *ListNode) {
	if root.Next.Next == nil {
		mid = root.Next
		root.Next = nil
		return root, mid, nil
	}
	fast, low, left, before := root, root, root, root
	for fast != nil && fast.Next != nil {
		before = low
		fast, low = fast.Next.Next, low.Next
	}
	before.Next = nil
	return left, low, low.Next
}

func main() {
	link5 := &ListNode{9, nil}
	link4 := &ListNode{5, link5}
	link3 := &ListNode{0, link4}
	link2 := &ListNode{-3, link3}
	link1 := &ListNode{-10, link2}
	bst := sortedListToBST(link1)
	print(bst)
}

func print(root *TreeNode) {
	if root != nil {
		print(root.Left)
		fmt.Printf("%d ", root.Val)
		print(root.Right)
	}
}
