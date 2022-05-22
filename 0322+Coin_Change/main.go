package main

import (
	"math"
)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
	}

	for _, c := range coins {
		for a := c; a <= amount; a++ {
			p := dp[a-c]
			if p != math.MaxInt {
				ans := 1 + p
				if dp[a] > ans {
					dp[a] = ans
				}
			}
		}
	}
	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}
