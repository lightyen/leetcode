package main

import (
	"app/util"
	"testing"
)

func TestCountSubstrings(t *testing.T) {
	util.Expect(t, countSubstrings("abc"), 3)
	util.Expect(t, countSubstrings("aba"), 4)
	util.Expect(t, countSubstrings("aaa"), 6)
	util.Expect(t, countSubstrings("aabaaa"), 12)
}
