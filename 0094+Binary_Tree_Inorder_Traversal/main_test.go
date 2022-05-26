package main

import (
	"app/util"
	"testing"
)

func TestCases(t *testing.T) {
	util.ExpectArray(t, inorderTraversal(&TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}},
	}), []int{1, 3, 2})
}
