package leetcode

import (
	"reflect"
	"unicode/utf8"
)

func checkInclusion(s1 string, s2 string) bool {
	s2Runes := []rune(s2)
	s1WithCounts := make(map[rune]int)
	for _, char := range s1 {
		s1WithCounts[char]++
	}

	left := 0
	state := make(map[rune]int)

	for right := 0; right < len(s2Runes); right++ {
		state[s2Runes[right]]++

		if right-left+1 == utf8.RuneCountInString(s1) {
			if reflect.DeepEqual(state, s1WithCounts) {
				return true
			}
		}

		if right-left+1 >= utf8.RuneCountInString(s1) {
			state[s2Runes[left]]--
			if state[s2Runes[left]] == 0 {
				delete(state, s2Runes[left])
			}

			left++
		}
	}

	return false
}
