package main

func reverseList(head *ListNode) *ListNode {
	_, reHead := reverseListWithHead(head)
	return reHead
}

func reverseListWithHead(head *ListNode) (next *ListNode, reHead *ListNode) {
	if head == nil || head.Next == nil {
		return head, head
	}
	node, reHead2 := reverseListWithHead(head.Next)
	node.Next = head
	head.Next = nil
	return head, reHead2
}
