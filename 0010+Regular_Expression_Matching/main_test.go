package main

import (
	"app/util"
	"testing"
)

func TestCases(t *testing.T) {
	util.Expect(t, isMatch("aab", "c*a*b"), true)
	util.Expect(t, isMatch("aa", "a"), false)
	util.Expect(t, isMatch("", ".*"), true)
	util.Expect(t, isMatch("", "."), false)
}
