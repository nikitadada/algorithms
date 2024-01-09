package sliding_window

func characterReplacement(s string, k int) int {
	left, res, maxFrequently := 0, 0, 0
	m := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		m[s[right]]++
		maxFrequently = max(maxFrequently, m[s[right]])

		if right-left+1-maxFrequently > k {
			m[s[left]]--
			left++
		}

		res = max(res, right-left+1)
	}

	return res
}
