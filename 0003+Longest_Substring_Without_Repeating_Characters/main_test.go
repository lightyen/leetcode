package main

import (
	"app/util"
	"testing"
)

func TestCases(t *testing.T) {
	util.Expect(t, lengthOfLongestSubstring("abcabcbb"), 3)
	util.Expect(t, lengthOfLongestSubstring("bbbbb"), 1)
	util.Expect(t, lengthOfLongestSubstring("pwwkew"), 3)
	util.Expect(t, lengthOfLongestSubstring("jxdlnaaij"), 6)
	util.Expect(t, lengthOfLongestSubstring("dvdf"), 3)
	util.Expect(t, lengthOfLongestSubstring("cwkcraoeqdqsxppgs"), 9)
}
