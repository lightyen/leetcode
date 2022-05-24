package main

// https://leetcode.com/problems/repeated-substring-pattern/

// KMP method

// 3 times:
//
// PPP
// PPPPPP
// xPPPPx

// 2 times:
//
// PP
// PPPP
// xPPx

// 1 time:
//
// P
// PP
// xx

func repeatedSubstringPattern(s string) bool {
	t := s + s
	return kmpFindIndex(t[1:len(t)-1], s) != -1
}

func buildTable(pat string) []int {
	table := make([]int, len(pat))
	table[0] = -1
	i, j := 2, 0
	for i < len(pat) {
		if pat[i-1] == pat[j] {
			j++
			table[i] = j
			i++
			continue
		}
		if j > 0 {
			j = table[j]
			continue
		}
		table[i] = 0
		i++
	}
	return table
}

func kmpFindIndex(s string, pat string) int {
	table := buildTable(pat)
	i, j := 0, 0
	for i < len(s) {
		if s[i] == pat[j] {
			j++
			if j == len(pat) {
				return i - j + 1
			}
			i++
			continue
		}
		if table[j] < 0 {
			j = 0
			i++
			continue
		}
		j = table[j]
	}
	return -1
}
