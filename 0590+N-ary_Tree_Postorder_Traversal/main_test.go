package main

import (
	"app/util"
	"testing"
)

func TestCases(t *testing.T) {
	util.ExpectArray(t, postorder(&Node{
		Val: 1,
		Children: []*Node{
			{Val: 3, Children: []*Node{
				{Val: 5},
				{Val: 6}}},
			{Val: 2},
			{Val: 4},
		},
	}), []int{5, 6, 3, 2, 4, 1})
}
