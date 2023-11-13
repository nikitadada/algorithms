package leetcode

import "sort"

func longestConsecutive2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	maxSeq := 1
	currentSeq := 1
	p1 := 0
	p2 := 1
	for p2 < len(nums) {
		if nums[p2]-1 == nums[p1] {
			currentSeq++

			p1 = p2
			p2++
		} else if nums[p1] == nums[p2] {
			p2++
			continue
		} else {
			currentSeq = 1

			p1 = p2
			p2++
		}
		maxSeq = max(currentSeq, maxSeq)
	}

	return maxSeq
}
