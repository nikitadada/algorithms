package sliding_window

func lengthOfLongestSubstring(s string) int {
	runes := []rune(s)

	state := make(map[rune]struct{})
	res, windowStart := 0, 0

	for windowEnd := 0; windowEnd < len(runes); windowEnd++ {
		for {
			_, ok := state[runes[windowEnd]]
			if !ok {
				break
			}

			delete(state, runes[windowStart])
			windowStart++
		}

		state[runes[windowEnd]] = struct{}{}

		res = max(res, windowEnd-windowStart+1)
	}

	return res
}
