package sliding_window

func lengthOfLongestSubstring(s string) int {
	runes := []rune(s)

	state := make(map[rune]struct{})
	res, left := 0, 0

	for i := 0; i < len(runes); i++ {
		for {
			_, ok := state[runes[i]]
			if !ok {
				break
			}

			delete(state, runes[left])
			left++
		}

		state[runes[i]] = struct{}{}

		res = max(res, i-left+1)
	}

	return res
}
