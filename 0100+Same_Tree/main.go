package main

import "app/util"

// https://leetcode.com/problems/same-tree/

type TreeNode = util.TreeNode[int]

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return p == nil && q == nil
}
