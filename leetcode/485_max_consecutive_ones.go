package leetcode

func findMaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	curCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			curCount++
			if curCount > maxCount {
				maxCount = curCount
			}
		} else {
			curCount = 0
		}
	}

	return maxCount
}
