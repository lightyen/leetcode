package main

import (
	"app/util"
	"testing"
)

func TestCases(t *testing.T) {
	util.Expect(t, findMedianSortedArrays([]int{1, 2}, []int{3, 4}), 2.5)
	util.Expect(t, findMedianSortedArrays([]int{1, 2}, []int{3, 4, 5}), 3)
	util.Expect(t, findMedianSortedArrays([]int{1, 3}, []int{2}), 2)
}
