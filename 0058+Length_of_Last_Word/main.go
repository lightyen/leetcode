package main

// https://leetcode.com/problems/length-of-last-word

func lengthOfLastWord(s string) int {
	i := len(s) - 1

	for i > 0 && s[i] == ' ' {
		i--
	}

	j := i
	for j >= 0 && s[j] != ' ' {
		j--
	}

	return i - j
}
