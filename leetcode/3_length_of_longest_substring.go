package leetcode

func lengthOfLongestSubstring2(s string) int {
	runes := []rune(s)

	if len(runes) == 1 {
		return 1
	}

	state := make(map[rune]int)

	res, left := 0, 0

	for right := 0; right < len(runes); right++ {
		if _, ok := state[runes[right]]; ok {
			state[runes[left]]--
			if state[runes[left]] == 0 {
				delete(state, runes[left])
			}

			left++
			right--
			continue
		} else {
			state[runes[right]]++
		}

		res = max(res, right-left+1)
	}

	return res
}
