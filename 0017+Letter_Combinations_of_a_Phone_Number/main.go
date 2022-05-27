package main

import "fmt"

// https://leetcode.com/problems/letter-combinations-of-a-phone-number/

var table = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	var result []string
	var fn func(int, string)
	fn = func(i int, s string) {
		if i == len(digits) {
			result = append(result, s)
			return
		}
		for _, c := range table[digits[i]] {
			fn(i+1, s+c)
		}
	}
	fn(0, "")
	fmt.Println(result)
	return result
}
