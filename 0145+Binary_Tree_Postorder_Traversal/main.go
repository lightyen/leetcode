package main

import (
	"app/util"
	"math"
)

// https://leetcode.com/problems/binary-tree-postorder-traversal/

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

func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node.Left != nil {
			traversal(node.Left)
		}
		if node.Right != nil {
			traversal(node.Right)
		}
		result = append(result, node.Val)
	}
	traversal(root)
	return result
}
