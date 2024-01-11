package sliding_window

func totalFruit(fruits []int) int {
	windowStart, res := 0, 0
	state := make(map[int]int, 3)

	for windowEnd := 0; windowEnd < len(fruits); windowEnd++ {
		state[fruits[windowEnd]]++

		for len(state) > 2 {
			state[fruits[windowStart]]--
			if state[fruits[windowStart]] == 0 {
				delete(state, fruits[windowStart])
			}

			windowStart++
		}

		res = Max(windowEnd-windowStart+1, res)
	}

	return res
}

func Max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}
