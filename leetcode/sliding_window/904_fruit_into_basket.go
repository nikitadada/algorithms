package sliding_window

func totalFruit(fruits []int) int {
	stateMap := make(map[int]int, 3)
	windowStart, res := 0, 0

	for windowEnd := 0; windowEnd < len(fruits); windowEnd++ {
		stateMap[fruits[windowEnd]]++

		if len(stateMap) > 2 {
			stateMap[fruits[windowStart]]--
			if stateMap[fruits[windowStart]] == 0 {
				delete(stateMap, fruits[windowStart])
			}

			windowStart++
			continue
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
