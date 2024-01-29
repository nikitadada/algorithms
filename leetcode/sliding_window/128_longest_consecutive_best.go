package sliding_window

func longestConsecutive(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	res := 0

	for _, num := range nums {
		m[num] = struct{}{}
	}

	for num := range m {
		if _, ok := m[num-1]; ok {
			continue
		}

		currElemOfConsecutive := num
		for {
			if _, ok := m[currElemOfConsecutive+1]; ok {
				currElemOfConsecutive++
			} else {
				break
			}
		}

		res = max(res, currElemOfConsecutive-num+1)
	}

	return res
}
