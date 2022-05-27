package main

import (
	"app/util"
	"sort"
	"testing"
)

func equal(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	i := map[string]bool{}
	j := map[string]bool{}
	var aKeys []string
	var bKeys []string
	for _, s := range a {
		i[s] = true
	}
	for _, s := range b {
		j[s] = true
	}
	for k := range i {
		aKeys = append(aKeys, k)
	}
	for k := range j {
		bKeys = append(bKeys, k)
	}
	sort.Strings(aKeys)
	sort.Strings(bKeys)
	for i := 0; i < len(aKeys); i++ {
		if aKeys[i] != bKeys[i] {
			return false
		}
	}
	return true
}

func TestCases(t *testing.T) {
	util.Expect(t, equal(generateParenthesis(4),
		[]string{"()()()()", "()()(())", "()(())()", "()(()())", "()((()))", "(())()()", "(())(())", "(()())()", "(()()())", "(()(()))", "((()))()", "((())())", "((()()))", "(((())))"}),
		true,
	)
}
