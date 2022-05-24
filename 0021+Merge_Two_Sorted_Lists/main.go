package main

// https://leetcode.com/problems/merge-two-sorted-lists

import "app/util"

type ListNode = util.ListNode[int]

func mergeTwoLists(a *ListNode, b *ListNode) *ListNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	var h *ListNode
	if a.Val < b.Val {
		h = a
		a = a.Next
	} else {
		h = b
		b = b.Next
	}

	p := h
	for a != nil && b != nil {
		if a.Val < b.Val {
			p.Next = a
			a = a.Next
		} else {
			p.Next = b
			b = b.Next
		}
		p = p.Next
	}
	if a != nil {
		p.Next = a
	} else {
		p.Next = b
	}
	return h
}
