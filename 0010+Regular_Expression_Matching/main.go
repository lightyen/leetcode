package main

// https://leetcode.com/problems/regular-expression-matching/

// related: #44

type Char struct {
	C     byte
	Multi bool
}

func parse(p string) []*Char {
	var s []*Char
	var t *Char
	for i := 0; i < len(p); i++ {
		if p[i] == '*' {
			t.Multi = true
			continue
		}
		if t != nil {
			s = append(s, t)
		}
		t = &Char{p[i], false}
	}
	s = append(s, t)
	return s
}

func isMatch(s string, p string) bool {
	pat := parse(p)
	dp := make([][]bool, len(pat)+1)
	for i := 0; i <= len(pat); i++ {
		dp[i] = make([]bool, len(s)+1)
	}

	dp[0][0] = true
	// for i := 0; i < len(s); i++ {
	// 	dp[0][i+1] = false
	// }

	for i := 0; i < len(pat); i++ {
		if dp[i][0] {
			if pat[i].Multi {
				dp[i+1][0] = true
			}
		}
	}

	for j := 0; j < len(s); j++ {
		for i := 0; i < len(pat); i++ {
			if pat[i].Multi {
				if pat[i].C == '.' {
					dp[i+1][j+1] = dp[i][j+1] || dp[i+1][j]
				} else {
					dp[i+1][j+1] = dp[i][j+1] || (dp[i+1][j] && s[j] == pat[i].C)
				}
			} else if pat[i].C == s[j] || pat[i].C == '.' {
				dp[i+1][j+1] = dp[i][j]
			}
		}
	}

	return dp[len(pat)][len(s)]
}
