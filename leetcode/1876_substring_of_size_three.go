package leetcode

func countGoodSubstrings(s string) int {
	runes := []rune(s)

	state := make(map[rune]int)
	leftPointer, res := 0, 0

	for i := 0; i < len(runes); i++ {
		state[runes[i]]++

		if i-leftPointer == 3 {
			leftCh := runes[leftPointer]

			state[leftCh]--
			if state[leftCh] == 0 {
				delete(state, leftCh)
			}
			leftPointer++
		}

		if len(state) == 3 {
			res++
		}
	}

	return res
}
