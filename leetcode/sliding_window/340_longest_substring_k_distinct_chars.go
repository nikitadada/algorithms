package sliding_window

func longestSubstringKDistinct(s string, k int) int {
	windowStart, res := 0, 0
	state := make(map[byte]int)

	for windowEnd := 0; windowEnd < len(s); windowEnd++ {
		state[s[windowEnd]]++

		if len(state) > k {
			state[s[windowStart]]--
			if state[s[windowStart]] == 0 {
				delete(state, s[windowStart])
			}

			windowStart++
			continue
		}

		res = max(res, windowEnd-windowStart+1)
	}

	return res
}
