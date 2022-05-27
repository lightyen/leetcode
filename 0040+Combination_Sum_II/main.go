package main

import (
	"sort"
)

// https://leetcode.com/problems/combination-sum-ii/

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	table := make([]int, 51)
	for _, v := range candidates {
		table[v]++
	}
	x := 0
	for i, v := range table {
		if v > 0 {
			candidates[x] = i
			x++
		}
	}
	candidates = candidates[:x]

	var result [][]int
	var fn func(int, int, []int, []int)
	fn = func(i int, sum int, table []int, nums []int) {
		if sum == target {
			result = append(result, nums)
			return
		}
		for j := i; j < len(candidates); j++ {
			c := candidates[j]
			if v := table[c]; v == 0 {
				continue
			}
			if sum+c > target {
				continue
			}
			arr := make([]int, len(nums)+1)
			copy(arr, nums)
			arr[len(nums)] = c
			t := make([]int, len(table))
			copy(t, table)
			t[c]--
			fn(j, sum+c, t, arr)
		}
	}
	fn(0, 0, table, nil)
	return result
}
