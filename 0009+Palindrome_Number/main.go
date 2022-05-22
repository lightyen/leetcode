package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	a := x
	b := 0
	for x != 0 {
		b = 10*b + x%10
		x /= 10
	}
	return a == b
}
