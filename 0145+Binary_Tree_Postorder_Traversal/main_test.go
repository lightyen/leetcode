package main

import (
	"app/util"
	"testing"
)

func TestCases(t *testing.T) {
	util.ExpectArray(t, postorderTraversal(&TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}},
	}), []int{3, 2, 1})
}
