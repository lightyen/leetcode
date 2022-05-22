package main

import (
	"app/util"
	"testing"
)

func TestCases(t *testing.T) {
	util.Expect(t, isValid("()[]{}"), true)
	util.Expect(t, isValid("([)]"), false)
	util.Expect(t, isValid("(]"), false)
	util.Expect(t, isValid("{[]}"), true)
}
