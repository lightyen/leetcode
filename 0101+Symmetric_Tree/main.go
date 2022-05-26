package main

import "app/util"

// https://leetcode.com/problems/symmetric-tree/

type TreeNode = util.TreeNode[int]

func isSymmetric(root *TreeNode) bool {
	var symme func(a *TreeNode, b *TreeNode) bool
	symme = func(a *TreeNode, b *TreeNode) bool {
		if a != nil && b != nil {
			return a.Val == b.Val && symme(a.Left, b.Right) && symme(a.Right, b.Left)
		}
		return a == nil && b == nil
	}
	return symme(root.Left, root.Right)
}
