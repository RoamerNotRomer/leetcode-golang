package sol

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre, cur, next *ListNode

	if head == nil {
		return head
	}

	cur = head
	next = head.Next

	for next != nil {
		cur.Next = pre
		pre = cur
		cur = next
		next = next.Next
	}

	cur.Next = pre

	return cur
}
