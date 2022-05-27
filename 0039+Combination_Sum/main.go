package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/combination-sum/

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var result [][]int
	var fn func(int, int, []int)
	fn = func(i int, sum int, nums []int) {
		if sum == target {
			result = append(result, nums)
			return
		}
		for j := i; j < len(candidates); j++ {
			c := candidates[j]
			if sum+c > target {
				continue
			}
			arr := make([]int, len(nums)+1)
			copy(arr, nums)
			arr[len(nums)] = c
			fn(j, sum+c, arr)
		}
	}
	fn(0, 0, nil)
	return result
}

func main() {
	fmt.Printf("%#v\n", combinationSum([]int{7}, 1))
}
