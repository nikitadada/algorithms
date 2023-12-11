package leetcode

func lengthOfLongestSubstring2(s string) int {
	runes := []rune(s)
	left, res := 0, 0

	state := make(map[rune]struct{})

	for right := 0; right < len(runes); right++ {
		if _, ok := state[runes[right]]; ok {
			delete(state, runes[left])
			left++
			right--
			continue
		} else {
			state[runes[right]] = struct{}{}
		}

		res = max(res, len(state))
	}

	return res
}
