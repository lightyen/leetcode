package main

import (
	"app/util"
	"testing"
)

func TestCases(t *testing.T) {
	util.Expect(t, isMatch("cdcb", "*c?b"), false)
	util.Expect(t, isMatch("abcabczzzde", "*abc???de*"), true)
	util.Expect(t, isMatch("aa", "a"), false)
	util.Expect(t, isMatch("aa", "*"), true)
	util.Expect(t, isMatch("aa", "b*"), false)
	util.Expect(t, isMatch("aa", "*b"), false)
	util.Expect(t, isMatch("aa", "a*"), true)
	util.Expect(t, isMatch("aa", "?a"), true)
	util.Expect(t, isMatch("xa", "?a"), true)
	util.Expect(t, isMatch("", "*****"), true)
	util.Expect(t, isMatch("", "*abc"), false)
	util.Expect(t, isMatch("bbbbbbbabbaabbabbbbaaabbabbabaaabbababbbabbbabaaabaab", "b*b*ab**ba*b**b***bba"), false)
	util.Expect(t, isMatch("bbaaaabaaaaabbabbabbabbababaabababaabbabaaabbaababababbabaabbabbbbbbaaaaaabaabbbbbabbbbabbabababaaaaa", "******aa*bbb*aa*a*bb*ab***bbba*a*babaab*b*aa*a****"), false)
}
