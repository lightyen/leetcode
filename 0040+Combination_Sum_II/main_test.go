package main

import (
	"app/util"
	"testing"
)

func TestCases(t *testing.T) {
	util.ExpectArrayArray(t, combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8), [][]int{
		{1, 1, 6},
		{1, 2, 5},
		{1, 7},
		{2, 6},
	})
}
