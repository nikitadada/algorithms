package leetcode

import "unicode"

func isPalindrome(s string) bool {
	runes := []rune(s)
	l := 0
	r := len(runes) - 1

	for l <= r {
		if !unicode.IsLetter(runes[l]) && !unicode.IsDigit(runes[l]) {
			l++
			continue
		}

		if !unicode.IsLetter(runes[r]) && !unicode.IsDigit(runes[r]) {
			r--
			continue
		}

		if unicode.ToLower(runes[l]) != unicode.ToLower(runes[r]) {
			return false
		}

		l++
		r--
	}

	return true
}
