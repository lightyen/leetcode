package main

func isValid(s string) bool {
	i := 0
	stack := make([]rune, len(s))
	table := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, c := range s {
		switch c {
		case '(':
			fallthrough
		case '[':
			fallthrough
		case '{':
			stack[i] = c
			i++
		case ')':
			fallthrough
		case ']':
			fallthrough
		case '}':
			if i <= 0 {
				return false
			}
			i--
			if stack[i] != table[c] {
				return false
			}
		}
	}
	return i == 0
}
