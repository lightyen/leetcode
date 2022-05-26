package main

import (
	"app/util"
	"math"
)

// https://leetcode.com/problems/validate-binary-search-tree/

type TreeNode = util.TreeNode[int]

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	val := math.MinInt
	isValid := true
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node.Left != nil {
			traversal(node.Left)
		}
		if node.Val <= val {
			isValid = false
			return
		}
		val = node.Val
		if node.Right != nil {
			traversal(node.Right)
		}
	}
	traversal(root)
	return isValid
}
