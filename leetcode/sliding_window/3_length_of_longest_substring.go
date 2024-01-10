package sliding_window

func lengthOfLongestSubstring2(s string) int {
	windowStart, res := 0, 0
	state := make(map[rune]struct{})
	runes := []rune(s)

	for windowEnd := 0; windowEnd < len(runes); windowEnd++ {
		if _, ok := state[runes[windowEnd]]; ok {
			delete(state, runes[windowStart])
			windowStart++
			windowEnd--
			continue
		} else {
			state[runes[windowEnd]] = struct{}{}
		}

		res = max(res, windowEnd-windowStart+1)
	}

	return res
}
