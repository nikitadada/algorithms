package leetcode

func lengthOfLastWord(s string) int {
	rightPointer := len(s) - 1

	lastLetterIndex := rightPointer
	allSpacesSkipped := false
	preLastSpaceIndex := 0
	for {
		if rightPointer < 0 {
			break
		}

		if s[rightPointer] != ' ' && !allSpacesSkipped {
			lastLetterIndex = rightPointer
			allSpacesSkipped = true
		}

		if allSpacesSkipped && s[rightPointer] == ' ' {
			preLastSpaceIndex = rightPointer + 1
			break
		}

		rightPointer--
	}

	return lastLetterIndex - preLastSpaceIndex + 1
}
