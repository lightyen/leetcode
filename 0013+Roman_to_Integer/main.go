package main

// https://leetcode.com/problems/roman-to-integer

func romanToInt(s string) int {
	m := 0
	result := 0
	tmp := 0
	for i := len(s) - 1; i >= 0; i-- {
		c := toInt(s[i])
		if c < m {
			result = result + tmp - c
			tmp = 0
			m = 0
		} else {
			tmp += c
			m = c
		}
	}
	return result + tmp
}

func toInt(c byte) int {
	switch c {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}
