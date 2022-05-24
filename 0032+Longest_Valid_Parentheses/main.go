package main

// https://leetcode.com/problems/longest-valid-parentheses

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func longestValidParentheses(s string) int {
	stack := make([]int, len(s)+1)
	si := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			stack[si] = i
			si++
		case ')':
			if si > 0 {
				if s[stack[si-1]] == '(' {
					si--
					continue
				}
			}
			stack[si] = i
			si++
		}
	}
	if si == 0 {
		return len(s)
	}

	result := stack[0]
	stack[si] = len(s)
	si++
	for i := 1; i < si; i++ {
		result = max(result, stack[i]-stack[i-1]-1)
	}
	return result
}
