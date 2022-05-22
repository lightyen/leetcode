package main

import (
	"app/util"
	"testing"
)

func TestCases(t *testing.T) {
	util.Expect(t, searchInsert([]int{1, 3, 5, 6}, 5), 2)
	util.Expect(t, searchInsert([]int{1, 3, 5, 6}, 2), 1)
	util.Expect(t, searchInsert([]int{1, 3, 5, 6}, 7), 4)
}
