package main

import (
	"app/util"
	"testing"
)

func TestCases(t *testing.T) {
	util.ExpectArrayArray(t, combinationSum([]int{2, 3, 6, 7}, 7), [][]int{{2, 2, 3}, {7}})
	util.ExpectArrayArray(t, combinationSum([]int{2, 3, 5}, 8), [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}})
}
