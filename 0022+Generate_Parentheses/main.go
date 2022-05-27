package main

// https://leetcode.com/problems/generate-parentheses/

func generateParenthesis(n int) []string {
	var result []string
	var fn func(int, int, string)
	m := n << 1
	fn = func(step, v int, s string) {
		if step == m {
			result = append(result, s)
			return
		}
		if v > 0 {
			fn(step+1, v-1, s+")")
		}
		if v < n && v < m-step {
			fn(step+1, v+1, s+"(")
		}
	}
	fn(0, 0, "")
	return result
}
