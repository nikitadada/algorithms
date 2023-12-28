package leetcode

func isAnagram(s string, t string) bool {
	m := make(map[rune]int)

	for _, b := range s {
		m[b]++
	}

	for _, b := range t {
		if _, ok := m[b]; !ok {
			return false
		} else {
			m[b]--
			if m[b] == 0 {
				delete(m, b)
			}
		}
	}

	if len(m) == 0 {
		return true
	}

	return false
}
