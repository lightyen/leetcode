package main

// https://leetcode.com/problems/median-of-two-sorted-arrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	x := 0
	y := 0
	push := func(value int) {
		x = y
		y = value
	}

	a := 0
	b := 0
	m := (len(nums1) + len(nums2)) >> 1
	for (a + b) <= m {
		exist1 := a < len(nums1)
		exist2 := b < len(nums2)
		if exist1 && exist2 {
			if nums1[a] < nums2[b] {
				push(nums1[a])
				a++
			} else {
				push(nums2[b])
				b++
			}
		} else if exist1 {
			push(nums1[a])
			a++
		} else {
			push(nums2[b])
			b++
		}
	}

	if (len(nums1)^len(nums2))&1 == 1 {
		return float64(y)
	}
	return float64(x+y) / 2.0
}
