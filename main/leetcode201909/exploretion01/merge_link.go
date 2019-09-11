package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil {
		l2.Next = mergeTwoLists(l1, l2.Next)
	} else if l2 == nil {
		l1.Next = mergeTwoLists(l1.Next, l2)
	} else {
		if l1.Val < l2.Val {
			l1.Next = mergeTwoLists(l1.Next, l2)
		} else {
			l2.Next = mergeTwoLists(l1, l2.Next)
		}
	}
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else {
		if l1.Val < l2.Val {
			return l1
		} else {
			return l2
		}
	}
}
