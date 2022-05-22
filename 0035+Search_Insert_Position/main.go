package main

func searchInsert(nums []int, target int) int {
	var search func(i, j int) int
	search = func(i, j int) int {
		if i == j {
			return i
		}
		m := (i + j) >> 1
		if target == nums[m] {
			return m
		}
		if target < nums[m] {
			return search(i, m)
		}
		return search(m+1, j)
	}
	return search(0, len(nums))
}
