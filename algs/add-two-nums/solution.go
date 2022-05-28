package algs

import (
	"github.com/yanghl22/playgo/structures"
)

type ListNode = structures.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	v1, v2, carry, current := 0, 0, 0, head

	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			v1 = 0
		} else {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			v2 = 0
		} else {
			v2 = l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: (v1 + v2 + carry) % 10}
		current = current.Next
		carry = (v1 + v2 + carry) / 10
	}
	return head.Next
}

func addTwoNumbers1(n1 *ListNode, n2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	v1, v2, carrier, current := 0, 0, 0, head

	for n1 != nil || n2 != nil {
		if n1 == nil {
			v1 = 0
		} else {
			v1 = n1.Val
			n1 = n1.Next
		}
		if n2 == nil {
			v2 = 0
		} else {
			v2 = n2.Val
			n2 = n2.Next
		}
		current.Val = (carrier + v1 + v2) % 10
		current.Next = &ListNode{Val: 0}
		carrier = (carrier + v1 + v2) / 10
	}
	current.Next = nil
	return head
}
