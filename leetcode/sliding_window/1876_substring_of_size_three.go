package sliding_window

func countGoodSubstrings(s string) int {
	runes := []rune(s)

	state := make(map[rune]int)
	windowStart, res := 0, 0

	for windowEnd := 0; windowEnd < len(runes); windowEnd++ {
		state[runes[windowEnd]]++

		if windowEnd-windowStart+1 > 3 {
			leftCh := runes[windowStart]

			state[leftCh]--
			if state[leftCh] == 0 {
				delete(state, leftCh)
			}
			windowStart++
		}

		if len(state) == 3 {
			res++
		}
	}

	return res
}
