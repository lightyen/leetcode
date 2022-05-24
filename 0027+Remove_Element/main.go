package main

func removeElement(nums []int, val int) int {
	t := 0
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if n != val {
			nums[t] = n
			t++
		}
	}
	return t
}
