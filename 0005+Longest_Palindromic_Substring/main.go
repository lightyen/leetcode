package main

// https://leetcode.com/problems/longest-palindromic-substring/

// babab
//
// b     a     b     a      b
// ba    ab    ba    ab
// bab   aba   bab
// baba  abab
// babab

func longestPalindrome(s string) string {
	a, b, c := 0, 1, 1

	dp := make([][]bool, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := 0; i < len(s); i++ {
		dp[0][i] = true
	}
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			dp[1][i] = true
			if c < 2 {
				c = 2
				a, b = i, i+c
			}
		}
	}

	for k := 3; k <= len(s); k++ {
		for i := 0; i+k <= len(s); i++ {
			if s[i] == s[i+k-1] {
				ok := dp[k-3][i+1]
				dp[k-1][i] = ok
				if ok && c < k {
					c = k
					a, b = i, i+c
				}
			}
		}
	}

	return s[a:b]
}
