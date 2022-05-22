package main

import (
	"app/util"
	"testing"
)

func TestCases(t *testing.T) {
	util.Expect(t, lengthOfLastWord("Hello World"), 5)
	util.Expect(t, lengthOfLastWord("HelloWorld"), 10)
	util.Expect(t, lengthOfLastWord("   fly me   to   the moon  "), 4)
	util.Expect(t, lengthOfLastWord("luffy is still joyboy"), 6)
}
