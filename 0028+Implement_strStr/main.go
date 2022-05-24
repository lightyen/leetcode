package main

// https://leetcode.com/problems/implement-strstr/

// KMP method

// func strStr(haystack string, needle string) int {
// 	if haystack == "" {
// 		return 0
// 	}
// 	for i := 0; i < len(haystack); i++ {
// 		match := true
// 		for j := 0; j < len(needle); j++ {
// 			if i+j >= len(haystack) {
// 				match = false
// 				break
// 			}
// 			if needle[j] != haystack[i+j] {
// 				match = false
// 				break
// 			}
// 		}
// 		if match {
// 			return i
// 		}
// 	}
// 	return -1
// }

func strStr(haystack string, needle string) int {
	return kmpFindIndex(haystack, needle)
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
