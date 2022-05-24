package main

// https://leetcode.com/problems/longest-substring-without-repeating-characters

// related: #53

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func lengthOfLongestSubstring(s string) int {
	a := 0
	b := 0
	maxLength := 0
	last := map[byte]int{}
	for ; b < len(s); b++ {
		char := s[b]
		if pos, ok := last[char]; ok {
			a = max(pos+1, a)
		}
		maxLength = max(maxLength, b-a+1)
		last[char] = b
	}
	return maxLength
}
