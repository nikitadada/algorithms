package leetcode

// Input: s = "ab#c", t = "ad#c"
// Output: true
// Explanation: Both s and t become "ac".
func backspaceCompare(s string, t string) bool {
	p1 := len(s) - 1
	p2 := len(t) - 1

	for p1 >= 0 || p2 >= 0 {
		p1 = getCurrentIndex(s, p1)
		p2 = getCurrentIndex(t, p2)

		if p1 < 0 && p2 < 0 {
			return true
		}
		if p1 < 0 || p2 < 0 {
			return false
		}
		if s[p1] != t[p2] {
			return false
		}

		p1--
		p2--
	}

	return true
}

func getCurrentIndex(s string, fromIndex int) int {
	toSkip := 0

	for fromIndex >= 0 {
		if s[fromIndex] == '#' {
			toSkip++
		} else if toSkip > 0 {
			toSkip--
		} else {
			break
		}

		fromIndex--
	}

	return fromIndex
}
