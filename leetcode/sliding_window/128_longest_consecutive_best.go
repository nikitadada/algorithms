package sliding_window

func longestConsecutive(nums []int) int {
	m := make(map[int]struct{}, len(nums))

	for _, v := range nums {
		m[v] = struct{}{}
	}

	resCount := 0

	for num := range m {
		if _, ok := m[num-1]; !ok {
			continue
		}

		cur := num
		for {
			if _, ok := m[cur+1]; !ok {
				break
			}
			cur++
		}

		resCount = max(resCount, cur)
	}

	return resCount
}
