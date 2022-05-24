package main

import (
	"app/util"
	"testing"
)

func except(nums []int, target int, excepted []int) bool {
	k := removeElement(nums, target)
	if k != len(excepted) {
		return false
	}
	for i := 0; i < k; i++ {
		if excepted[i] != nums[i] {
			return false
		}
	}
	return true
}

func TestCases(t *testing.T) {
	util.Expect(t, except([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, []int{0, 1, 3, 0, 4}), true)
}
