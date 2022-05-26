package main

import (
	"app/util"
)

// https://leetcode.com/problems/binary-tree-inorder-traversal/

type TreeNode = util.TreeNode[int]

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node.Left != nil {
			traversal(node.Left)
		}
		result = append(result, node.Val)

		if node.Right != nil {
			traversal(node.Right)
		}
	}
	traversal(root)
	return result
}
