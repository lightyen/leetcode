package main

//    a  a  b  a  a  a
// a [0, 0, 0, 0, 0, 1]
// a [1, 0, 0, 1 ,1]
// a [0, 1, 0, 1]
// b [0, 0, 1]
// a [1, 1]
// a [1]

func countSubstrings(s string) int {
	dp := make([]int, len(s))
	dp[0] = 1
	ans := 1
	for b := 1; b < len(s); b++ {
		for a := 0; a < b; a++ {
			if s[a] != s[b] {
				dp[a] = 0
				continue
			}
			if b-a <= 2 || dp[a+1] != 0 {
				dp[a] = 1
				ans++
				continue
			}
			dp[a] = 0
		}
		ans++
		dp[b] = 1
	}
	return ans
}
