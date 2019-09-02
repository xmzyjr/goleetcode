package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	temp := swapPairs(head.Next.Next)
	temp1 := head.Next
	head.Next = temp
	temp1.Next = head
	return temp1
}
