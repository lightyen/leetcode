package main

import (
	"app/util"
	"testing"
)

func TestCases(t *testing.T) {
	util.Expect(t, isValidBST(&TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 4},
		Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 7}},
	}), false)
}
