package main

import (
	"app/util"
	"testing"
)

func TestCases(t *testing.T) {
	util.Expect(t, repeatedSubstringPattern("aabc"), false)
	util.Expect(t, repeatedSubstringPattern("aaa"), true)
	util.Expect(t, repeatedSubstringPattern("abaaabaa"), true)
	util.Expect(t, repeatedSubstringPattern("abcabcabc"), true)
	util.Expect(t, repeatedSubstringPattern("ababab"), true)
	util.Expect(t, repeatedSubstringPattern("abcabc"), true)
	util.Expect(t, repeatedSubstringPattern("aabbcc"), false)
	util.Expect(t, repeatedSubstringPattern("abcdefabcdef"), true)
}
