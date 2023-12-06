package leetcode

func totalFruit(fruits []int) int {
	stateMap := make(map[int]int, 3)
	left := 0

	countFruits := 0
	for right := 0; right < len(fruits); right++ {
		stateMap[fruits[right]]++

		for len(stateMap) > 2 {
			stateMap[fruits[left]]--
			if stateMap[fruits[left]] == 0 {
				delete(stateMap, fruits[left])
			}

			left++
		}

		countFruits = max(right-left+1, countFruits)
	}

	return countFruits
}
