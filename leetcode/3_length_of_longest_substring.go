package leetcode

func lengthOfLongestSubstring2(s string) int {
	runes := []rune(s)

	if len(runes) == 1 {
		return 1
	}

	state := make(map[rune]int)

	res, left := 0, 0

	for i := 0; i < len(runes); i++ {
		if _, ok := state[runes[i]]; ok {
			leftCh := runes[left]
			state[leftCh]--
			if state[leftCh] == 0 {
				delete(state, leftCh)
			}

			res = max(res, i-left)
			left++
			i--
		} else {
			state[runes[i]]++
			res = max(res, i+1-left)
		}

		res = max(res, i-left)
	}

	return res
}
