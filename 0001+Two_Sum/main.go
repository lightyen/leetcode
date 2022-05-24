package main

// https://leetcode.com/problems/two-sum

func twoSum(nums []int, target int) []int {
	dp := map[int]int{}
	for j, num := range nums {
		if i, ok := dp[num]; ok {
			return []int{i, j}
		} else {
			dp[target-num] = j
		}
	}
	return nil
}
