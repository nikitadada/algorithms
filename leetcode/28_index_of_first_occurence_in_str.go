package leetcode

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	p1 := 0
	p2 := 0

	runesS1 := []rune(haystack)
	runesS2 := []rune(needle)

	if len(runesS2) > len(runesS1) {
		return -1
	}

	for p1 < len(runesS1) && p2 < len(runesS2) {
		if runesS1[p1] == runesS2[p2] {
			p1++
			p2++
			if p2 == len(runesS2) {
				return p1 - len(runesS2)
			}
		} else {
			p1 = p1 - p2 + 1
			p2 = 0
		}
	}

	return -1
}
