package main

import (
	"strings"
)

// https://leetcode.com/problems/wildcard-matching/

// related: #10

func reduce(s string) string {
	b := &strings.Builder{}
	for i := 0; i < len(s); i++ {
		b.WriteByte(s[i])
		if s[i] == '*' {
			for j := i + 1; j < len(s); j++ {
				if s[j] == '*' {
					i = j
				} else {
					break
				}
			}
		}
	}
	return b.String()
}

func isMatch(s string, p string) bool {
	p = reduce(p)

	dp := make([][]bool, len(p)+1)
	for i := 0; i <= len(p); i++ {
		dp[i] = make([]bool, len(s)+1)
	}

	dp[0][0] = true
	// for i := 0; i < len(s); i++ {
	// 	dp[0][i+1] = false
	// }

	for i := 0; i < len(p); i++ {
		if dp[i][0] {
			if p[i] == '*' {
				dp[i+1][0] = true
			}
		}
	}

	for j := 0; j < len(s); j++ {
		for i := 0; i < len(p); i++ {
			if p[i] == '*' {
				dp[i+1][j+1] = dp[i][j+1] || dp[i+1][j]
			} else if p[i] == s[j] || p[i] == '?' {
				dp[i+1][j+1] = dp[i][j]
			}
		}
	}

	return dp[len(p)][len(s)]
}
