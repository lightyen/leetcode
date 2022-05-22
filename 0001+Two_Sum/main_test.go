package main

import (
	"app/util"
	"testing"
)

func TestCases(t *testing.T) {
	util.ExpectArray(t, twoSum([]int{3, 2, 4}, 6), []int{1, 2})
	util.ExpectArray(t, twoSum([]int{3, 3}, 6), []int{0, 1})
}
