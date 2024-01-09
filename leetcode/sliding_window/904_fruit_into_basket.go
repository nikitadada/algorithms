package sliding_window

func totalFruit(fruits []int) int {
	stateMap := make(map[int]int, 3)
	left, res := 0, 0

	for right := 0; right < len(fruits); right++ {
		stateMap[fruits[right]]++

		if len(stateMap) > 2 {
			stateMap[fruits[left]]++

		}
	}

	return res
}
