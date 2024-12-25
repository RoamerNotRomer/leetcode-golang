package sol

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var pre, hair, next, tail *ListNode
	if head == nil {
		return head
	}
	hair = &ListNode{Val: 0}
	hair.Next = head
	pre = hair

	for head != nil {
		tail = pre
		i := 0
		for i < k {
			i++
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		next = tail.Next
		//tail.Next = nil
		head, tail = reverseList(head, tail)
		pre.Next = head
		tail.Next = next
		pre = tail
		head = tail.Next
	}

	return hair.Next
}

func reverseList(head, tail *ListNode) (*ListNode, *ListNode) {
	var prev, p *ListNode

	prev = tail.Next
	p = head

	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}

	return tail, head
}
