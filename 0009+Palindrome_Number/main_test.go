package main

import (
	"app/util"
	"testing"
)

func TestCases(t *testing.T) {
	util.Expect(t, isPalindrome(129921), true)
	util.Expect(t, isPalindrome(12345), false)
	util.Expect(t, isPalindrome(12521), true)
}
