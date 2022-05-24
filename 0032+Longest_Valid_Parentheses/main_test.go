package main

import (
	"app/util"
	"testing"
)

func TestCases(t *testing.T) {
	util.Expect(t, longestValidParentheses("(()"), 2)
	util.Expect(t, longestValidParentheses("(((()"), 2)
	util.Expect(t, longestValidParentheses("()()"), 4)
	util.Expect(t, longestValidParentheses("())"), 2)
	util.Expect(t, longestValidParentheses("((())))(())()())("), 8)
	util.Expect(t, longestValidParentheses(""), 0)
	util.Expect(t, longestValidParentheses("(()()()"), 6)
	util.Expect(t, longestValidParentheses("(()()())"), 8)
	util.Expect(t, longestValidParentheses("(())(()()()"), 6)
}
