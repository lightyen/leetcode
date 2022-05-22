package main

import (
	"app/util"
	"testing"
)

func TestCases(t *testing.T) {
	util.Expect(t, romanToInt("III"), 3)
	util.Expect(t, romanToInt("LVIII"), 58)
	util.Expect(t, romanToInt("MCMXCIV"), 1994)
}
