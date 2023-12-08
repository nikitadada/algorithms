package leetcode

import (
	"reflect"
	"unicode/utf8"
)

func findAnagrams(s string, p string) []int {
	sRunes := []rune(s)

	pWithCounts := make(map[rune]int)
	for _, char := range p {
		pWithCounts[char]++
	}

	left := 0
	windowState := make(map[rune]int)
	var res []int

	for right := 0; right < len(sRunes); right++ {
		windowState[sRunes[right]]++

		if right-left+1 == utf8.RuneCountInString(p) {
			if reflect.DeepEqual(windowState, pWithCounts) {
				res = append(res, left)
			}
		}

		if right-left+1 >= utf8.RuneCountInString(p) {
			windowState[sRunes[left]]--
			if windowState[sRunes[left]] == 0 {
				delete(windowState, sRunes[left])
			}

			left++
		}
	}

	return res
}
