package main

import (
	"app/util"
	"testing"
)

func TestCases(t *testing.T) {
	util.Expect(t, coinChange([]int{2, 5, 10, 1}, 27), 4)
	util.Expect(t, coinChange([]int{2}, 3), -1)
	util.Expect(t, coinChange([]int{1}, 0), 0)
	util.Expect(t, coinChange([]int{186, 419, 83, 408}, 6249), 20)
}
