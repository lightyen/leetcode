package main

import "app/util"

type ListNode = util.ListNode[int]

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var tmp = &ListNode{}
	var q = 0
	var p = tmp
	for l1 != nil || l2 != nil {
		sum := q
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		q = sum / 10
		p.Next = &ListNode{
			Val: sum % 10,
		}
		p = p.Next
	}
	if q > 0 {
		p.Next = &ListNode{
			Val: q,
		}
	}
	return tmp.Next
}
