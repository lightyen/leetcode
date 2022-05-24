package main

// https://leetcode.com/problems/maximum-subarray

// related: #3

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxSubArray(nums []int) int {
	maxVal := nums[0]
	cur := nums[0]
	for _, n := range nums[1:] {
		cur = max(n, cur+n)
		maxVal = max(maxVal, cur)
	}
	return maxVal
}
