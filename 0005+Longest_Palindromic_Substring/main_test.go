package main

import (
	"app/util"
	"testing"
)

func TestCases(t *testing.T) {
	util.Expect(t, longestPalindrome("babad"), "bab")
}
