package main

import (
	"app/util"
)

// https://leetcode.com/problems/binary-tree-preorder-traversal/

type TreeNode = util.TreeNode[int]

func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		result = append(result, node.Val)
		if node.Left != nil {
			traversal(node.Left)
		}
		if node.Right != nil {
			traversal(node.Right)
		}
	}
	traversal(root)
	return result
}
